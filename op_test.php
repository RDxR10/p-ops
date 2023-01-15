<?php
class Polynomial
{
  private $coefficients;
  public function __construct(array $coefficients)
  {
    $this->coefficients = $coefficients;
  }

  public function add(Polynomial $other)
  {
    $length = max(count($this->coefficients), count($other->coefficients));
    $result = [];
    for ($i = 0; $i < $length; $i++) {
      $result[$i] = ($this->coefficients[$i] ?? 0) + ($other->coefficients[$i] ?? 0);
    }

    return new Polynomial($result);
  }
  public function subtract(Polynomial $other)
  {
    $length = max(count($this->coefficients), count($other->coefficients));
    $result = [];
    for ($i = 0; $i < $length; $i++) {
      $result[$i] = ($this->coefficients[$i] ?? 0) - ($other->coefficients[$i] ?? 0);
    }
    
    return new Polynomial($result);
  }
  public function multiply(Polynomial $other)
  {

    $length = count($this->coefficients) + count($other->coefficients) - 1;
    $result = array_fill(0, $length, 0);
    for ($i = 0; $i < count($this->coefficients); $i++) {
      for ($j = 0; $j < count($other->coefficients); $j++) {
        $result[$i + $j] += $this->coefficients[$i] * $other->coefficients[$j];
      }
    }

    return new Polynomial($result);
  }

  public function square()
  {
    return $this->multiply($this);
  }
  public function evaluate($x)
  {
    $result = 0;
    for ($i = count($this->coefficients) - 1; $i >= 0; $i--) {
      $result = $result * $x + $this->coefficients[$i];
    }

    return $result;
  }
  public function __toString()
  {
    $result = '';

    for ($i = count($this->coefficients) - 1; $i >= 0; $i--) {
      if ($this->coefficients[$i] == 0) {
        continue;
      }
      $result .= $this->coefficients[$i] . 'x^' . $i . ' + ';
    }

    return rtrim($result, ' + ');
  }
}
