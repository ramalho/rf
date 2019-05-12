#! /usr/bin/env elixir

defmodule RuneFinder do
  @db_path Path.join(__DIR__, "UnicodeData.txt")

  defp display(code, name) do
    rune = <<String.to_integer(code, 16)::utf8>>
    IO.puts("U+#{code}\t#{rune}\t#{name}")
  end

  defp find(query_words) do
    count =
      @db_path
      |> File.stream!()
      |> Enum.reduce(0, fn line, count ->
        [code, name | _] = String.split(line, ";")

        if MapSet.subset?(query_words, tokenize(name)) do
          display(code, name)
          count + 1
        else
          count
        end
      end)

    IO.puts("(#{count} found)")
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
