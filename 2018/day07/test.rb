require 'pqueue'

data = File.read('data.txt').chomp.split("\n").map{ |d| d.match(/Step ([A-Z])[a-z ]+([A-Z])/)[1..2] }

class Node
  attr_accessor :val, :children

  def initialize(val)
    @val = val
    @children = []
  end
end

nodes = []
data.each do |before, after|
  added = nodes.map(&:val)
  nodes << Node.new(before) unless added.include? before
  nodes << Node.new(after) unless added.include? after

  before = nodes.find { |n| n.val == before }
  after = nodes.find { |n| n.val == after }

  before.children << after unless before.children.include? after
end

current_seconds = 0
base_seconds = 60

workers = { # value structure: [working_node, second_total_when_finished]
  1 => nil,
  2 => nil,
  3 => nil,
  4 => nil,
  5 => nil
}

until nodes.empty?
  # check for finished jobs
  workers.reject { |_, job| job.nil? }.each do |id, job|
    if job[1] == current_seconds
      workers[id] = nil
      nodes.delete(job[0])
    end
  end

  # get new jobs
  workers.select { |_, job| job.nil? }.each do |id, _|
    work_node = nodes.find { |n| nodes.none? { |n2| n2.children.include? n } && !workers.values.compact.map { |w| w[0] }.include?(n) }
    break unless work_node

    finish_time = current_seconds + work_node.val.ord - 64 + base_seconds
    workers[id] = [work_node, finish_time]
  end

  current_seconds += 1
end

puts current_seconds - 1
