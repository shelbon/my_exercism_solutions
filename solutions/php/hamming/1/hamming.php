<?php


function distance(string $strandA, string $strandB): int
{
    $hammingDistance = 0;
    if (strlen($strandA) != strlen($strandB)) {
        throw new InvalidArgumentException("DNA strands must be of equal length.");
    }
    for ($i = 0; $i < strlen($strandA); $i += 1) {
        if ($strandA[$i] != $strandB[$i]) {
            $hammingDistance++;
        }
    }
    return $hammingDistance;
}
