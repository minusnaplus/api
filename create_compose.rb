require 'erb'
require 'optparse'

# Define default variable values
defaults = {
  image_name: 'nginx',
  image_tag_1: 'proxy',
  image_tag_2: 'fastify',
  image_tag_3: 'fiber',
  environment: 'production',
  database_host: 'db',
  memory_limit: '512M'
}

OptionParser.new do |opts|
  opts.banner = 'Usage: generate-docker-compose.rb [options]'

  opts.on('-n', '--name NAME', 'Docker image name') do |name|
    defaults[:image_name] = name
  end

  opts.on('-t', '--tag TAG', 'Docker image tag') do |tag|
    defaults[:image_tag] = tag
  end

  opts.on('-e', '--environment ENV', 'Application environment') do |env|
    defaults[:environment] = env
  end

  opts.on('-d', '--database HOST', 'Database host') do |host|
    defaults[:database_host] = host
  end

  opts.on('-m', '--memory LIMIT', 'Container memory limit') do |limit|
    defaults[:memory_limit] = limit
  end
end.parse!

template_path = File.join(__dir__, 'compose.yml.erb')

template = File.read(template_path)
variables = defaults
context = binding

result = ERB.new(template).result(context)

puts result