class Junction
  attr_accessor :x, :y, :z, :circuit

  def initialize(x, y, z)
    @x = x
    @y = y
    @z = z
    @circuit = Circuit.new
    @circuit.add_juction!(self)
  end

  def to_s
    "(#{@x}, #{@y}, #{@z})"
  end

  # Straight Line Distance
  def sld(other)
    ((@x - other.x)**2 + (@y - other.y)**2 + (@z - other.z)**2)**0.5
  end
end

class Circuit
  attr_reader :junctions

  def initialize
    @junctions = Set.new
  end

  def add_juction!(junction)
    @junctions.add(junction)
    junction.circuit = self

    self
  end

  def to_s
    junctions_str = @junctions.map(&:to_s).join(', ')

    "Circuit(#{junctions_str})"
  end
end
