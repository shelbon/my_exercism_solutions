defmodule RomanNumerals do
    @roman_numerals [
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"}
  ]
 @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) do
    numeral_helper(number, "")
  end

  defp numeral_helper(0, result), do: result

  defp numeral_helper(number, result) do
    {value, symbol} =
      Enum.find(@roman_numerals, fn {value, _symbol} ->
        value <= number
      end)

    numeral_helper(number - value, result <> symbol)
  end
end 
