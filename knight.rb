Bundler.require

BOARD_SIZE = 8

class Placement
  attr_reader :row, :column

  def initialize(row:, column:)
    @row    = row
    @column = column
  end

  def inspect
    "(#{@row+1},#{@column+1})"
  end

  def ==(other)
    @row == other.row &&
      @column == other.column
  end

  def self.from_index(i)
    Placement.new(row: i/BOARD_SIZE, column: i%BOARD_SIZE)
  end

  def possibilities(visited)
    [
      [-2,-1],
      [-2,+1],
      [+2,-1],
      [+2,+1],
      [-1,-2],
      [-1,+2],
      [+1,-2],
      [+1,+2],
    ].map do |delta|
      Placement.new(
        row:    @row+delta[0],
        column: @column+delta[1],
      ) if (0..BOARD_SIZE-1).include?(@row+delta[0]) && (0..BOARD_SIZE-1).include?(@column+delta[1]) && !visited.include?(Placement.new(row: @row+delta[0], column: @column+delta[1]))
    end.compact
  end
end

def place_rest(visited)
  return visited if visited.length == BOARD_SIZE*BOARD_SIZE

  current_placement = visited.last

  possibilities = current_placement.possibilities(visited)
  if possibilities.nil? || possibilities.empty?
    puts "#{visited.inspect} cannot progress"
    return nil
  end

  current_placement.possibilities(visited).each do |p|
    if !visited.include?(p)
      updated_placements = visited + [p]
      fully_placed = place_rest(updated_placements)
      return fully_placed if !fully_placed.nil?
    end
  end

  return nil
end

current_placements = [Placement.new(row: 0, column: 0)]

final_placements = place_rest(current_placements)
debugger

pp current_placement
