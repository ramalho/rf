#! /usr/bin/env elixir

defmodule RuneFinder do
  defp display(code, name) do
    rune = <<String.to_integer(code, 16)::utf8>>
    IO.puts("U+#{code}\t#{rune}\t#{name}")
  end

  defp select(line_stream, query_words) do
    Enum.reduce(line_stream, 0, fn line, count ->
      [code, name | _] = String.split(line, ";")

      if MapSet.subset?(query_words, tokenize(name)) do
        display(code, name)
        count + 1
      else
        count
      end
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
