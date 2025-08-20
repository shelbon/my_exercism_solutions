defmodule FreelancerRates do
 @billable_days 22
  def daily_rate(hourly_rate) do
    8.0*hourly_rate
  end

  def apply_discount(before_discount, discount) do
    (1-discount/100)*before_discount
  end

  def monthly_rate(hourly_rate, discount) do
     raw_rate=apply_discount(daily_rate(hourly_rate),discount)* @billable_days
     trunc(ceil(raw_rate))
  end

  def days_in_budget(budget, hourly_rate, discount) do
      days= budget/apply_discount(daily_rate(hourly_rate),discount)
      Float.floor(days,1)
  end
end
