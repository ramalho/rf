#! /usr/bin/env elixir

defmodule RuneFinder do
  defp display(code_str, name) do
    code = String.to_integer(code_str, 16)
    char = List.to_string([code])
    IO.puts("U+#{code_str}\t#{char}\t#{name}")
  end

  defp select(lines, query_words) do
    Enum.count(lines, fn line ->
      [code_str, name | _] = String.split(line, ";")
      if MapSet.subset?(query_words, tokenize(name)), do: display(code_str, name)
    end) 
  end

  defp summary(count), do: IO.puts("(#{count} found)")

  defp find(query_words) do
    Path.join(__DIR__, "UnicodeData.txt")
    |> File.stream!()
    |> select(query_words)
    |> summary
  end

  defp tokenize(text) do
    text
    |> String.replace("-", " ")
    |> String.split()
    |> MapSet.new()
  end

  def main([]), do: IO.puts("Please provide words to find.")

  def main(args) do
    args
    |> Enum.join(" ")
    |> String.upcase()
    |> tokenize
    |> find
  end
end

RuneFinder.main(System.argv())
