#! /usr/bin/env elixir

defmodule RuneFinder do
  defp parse(line) do
    [code, name | _] = String.split(line, ";")
    {code, name, tokenize(name)}
  end

  defp find(query_words) do
    count =
      Path.join(__DIR__, "UnicodeData.txt")
      |> File.stream!()
      |> Enum.reduce(0, fn line, count ->
        {code, name, name_words} = parse(line)

        if MapSet.subset?(query_words, name_words) do
          rune = <<String.to_integer(code, 16)::utf8>>
          IO.puts("U+#{code}\t#{rune}\t#{name}")
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
