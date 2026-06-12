defmodule KitchenCalculator do
  def get_volume(volume_pair) do
    # Please implement the get_volume/1 function
    {_,volume}=volume_pair
    volume
  end

def to_milliliter({:cup, value}), do: {:milliliter, value * 240}
def to_milliliter({:milliliter, value}), do: {:milliliter, value}
def to_milliliter({:fluid_ounce, value}), do: {:milliliter, value * 30}
def to_milliliter({:teaspoon, value}), do: {:milliliter, value * 5}
def to_milliliter({:tablespoon, value}), do: {:milliliter, value * 15}
def to_milliliter(_), do: {:error, :unknown_unit}

def from_milliliter({:milliliter, value}, :milliliter), do: {:milliliter, value}
def from_milliliter({:milliliter, value}, :cup), do: {:cup, value / 240}
def from_milliliter({:milliliter, value}, :fluid_ounce), do: {:fluid_ounce, value / 30}
def from_milliliter({:milliliter, value}, :teaspoon), do: {:teaspoon, value / 5}
def from_milliliter({:milliliter, value}, :tablespoon), do: {:tablespoon, value / 15}

def from_milliliter({:milliliter, _value}, _unit), do: {:error, :unknown_unit}
def from_milliliter(_volume_pair, _unit), do: {:error, :expected_milliliter}

def convert(volume_pair, unit) do
  milliliter_pair = to_milliliter(volume_pair)
  from_milliliter(milliliter_pair, unit)
end
end
