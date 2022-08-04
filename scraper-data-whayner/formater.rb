#
# Simple script to remove some unwanted characters
#

# Receive a file 
text = File.read(ARGV[0])
text = text.gsub(/\\ +|\\n|\\\\n *| */,' ')
text = text.gsub(/  +/,' ')

# save as file.formated
File.open(ARGV[0]+".formated", "w") {|file| file.puts text}

exit

