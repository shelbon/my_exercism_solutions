<?php
function from($original_date)
{
    $date=clone $original_date;
  return   $date->add(new DateInterval("PT1000000000S"));
}