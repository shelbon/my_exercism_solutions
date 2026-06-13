defmodule HighSchoolSweetheart do
  def first_letter(name) do
     String.trim(name) 
       |> String.first()
  end

  def initial(name) do
   first_letter(name)
     |> String.upcase() 
     |> Kernel.<>(".")
  end

  def initials(full_name) do
      String.split(full_name)
      |> case do
         [] -> ""
         [first_name] -> initial(first_name)
         [first_name, last_name | _rest] -> initial(first_name) <> " " <> initial(last_name)
      end
  end

 def pair(full_name1, full_name2) do
  line = "❤-------------------❤"
  "#{line}\n|  #{initials(full_name1)}  +  #{initials(full_name2)}  |\n#{line}\n"
end
end
