require 'erb'

# Define the variables that will be used in the ERB template
containers = [
  {
    name: "container1",
    image: "image1",
    env: { VAR1: "value1", VAR2: "value2" },
    ports: ["8000:80", "8443:443"]
  },
  {
    name: "container2",
    image: "image2",
    env: { VAR3: "value3", VAR4: "value4" },
    ports: ["9000:80", "9443:443"]
  }
]

# Load the ERB template from file
template = File.read('docker-compose.yml.erb')

# Create a new ERB object based on the template
renderer = ERB.new(template)

# Render the ERB template with the given variables
result = renderer.result(binding)

# Output the result to STDOUT
puts result
