defmodule Secrets do
  import Bitwise

  def secret_add(secret) do
    fn value -> add(value, secret) end
  end

  defp add(current, 0), do: current

  defp add(current, carry) do
    new_current = bxor(current, carry)
    new_carry = bsl(band(current, carry), 1)

    add(new_current, new_carry)
  end

  def secret_subtract(secret) do
    fn value -> subtract(value, secret) end
  end

  defp subtract(value, secret) when value >= secret do
    subtract_positive_result(value, secret)
  end

  defp subtract(value, secret) do
    magnitude = subtract_positive_result(secret, value)
    negate(magnitude)
  end

  defp subtract_positive_result(a, 0), do: a

  defp subtract_positive_result(a, b) do
    new_a = bxor(a, b)
    new_b = bsl(band(bnot(a), b), 1)

    subtract_positive_result(new_a, new_b)
  end

  defp negate(value) do
    add(bnot(value), 1)
  end

  def secret_multiply(secret) do
    fn value -> multiply(value, secret, 0) end
  end

  defp multiply(_current_value, 0, result), do: result

  defp multiply(current_value, remaining_secret, result) do
    new_result =
      if last_bit_is_set?(remaining_secret) do
        add(result, current_value)
      else
        result
      end

    new_current_value = bsl(current_value, 1)
    new_remaining_secret = bsr(remaining_secret, 1)

    multiply(new_current_value, new_remaining_secret, new_result)
  end

  defp last_bit_is_set?(value) do
    band(value, 1) == 1
  end

  def secret_divide(secret) do
    fn value -> divide(value, secret, 0) end
  end

  defp divide(remainder, divisor, quotient) when remainder < divisor do
    quotient
  end

  defp divide(remainder, divisor, quotient) do
    new_remainder = subtract(remainder, divisor)
    new_quotient = add(quotient, 1)

    divide(new_remainder, divisor, new_quotient)
  end

  def secret_and(secret) do
    fn value -> band(value, secret) end
  end

  def secret_xor(secret) do
    fn value -> bxor(value, secret) end
  end

  def secret_combine(f1, f2) do
    fn value -> f2.(f1.(value)) end
  end
end