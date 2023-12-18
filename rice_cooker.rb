class RiceCooker
    def initialize
      @power_connected = false
      @lid_open = false
      @food_added = false
      @water_added = false
      @lid_closed = false
      @cooking = false
      @max_cup_limit = 5
      @food_quantity = 0
      @water_quantity = 0
    end
  
    def display_menu
      puts "Here are the steps to follow:"
      puts "1. Connect the power"
      puts "2. Open the lid"
      puts "3. Pour the food"
      puts "4. Pour water"
      puts "5. Close the lid"
      puts "6. Set the button to cook"
      puts "7. Cook"
      puts "8. Disconnect the power"
      puts "0. Quit"
    end
  
    def handle_user_choice(choice)
      case choice
      when 1 then connect_power
      when 2 then open_lid
      when 3 then pour_food
      when 4 then pour_water
      when 5 then close_lid
      when 6 then start_cooking
      when 7 then cook
      when 8 then disconnect_power
      when 0 then exit_program
      else
        puts "Invalid option. Please choose again."
      end
    end
  
    def connect_power
      if !@power_connected
        @power_connected = true
        puts "The power is connected, the current is flowing"
        puts "The warm button lights up"
      else
        puts "The power is already connected"
      end
    end
  
    def open_lid
      if !@lid_open
        @lid_open = true
        puts "The lid is open"
      else
        puts "The lid is already open"
      end
    end
  
    def pour_food
      if @lid_open
        puts "Quantity of food (in cups): "
        food_quantity = gets.to_i
        if food_quantity > 0 && food_quantity <= @max_cup_limit
          @food_added = true
          @food_quantity += food_quantity
          puts "Food added successfully"
        else
          puts "Invalid quantity of food"
        end
      else
        puts "The lid must be open to pour the food"
      end
    end
  
    def pour_water
      if @lid_open && @food_added
        puts "Quantity of water (in cups): "
        water_quantity = gets.to_i
        if water_quantity > 0 && water_quantity <= @max_cup_limit
          @water_added = true
          @water_quantity += water_quantity
          puts "Water added successfully"
        else
          puts "Invalid quantity of water"
        end
      else
        puts "The lid must be open, and food must be added before pouring water"
      end
    end
  
    def close_lid
      if @lid_open
        @lid_closed = true
        puts "The lid is closed"
      else
        puts "The lid is already closed"
      end
    end
  
    def start_cooking
      if @lid_closed
        @cooking = true
        puts "Cooking has started"
      else
        puts "The lid must be closed to start cooking"
      end
    end
  
    def cook
      if @cooking && @water_quantity >= @food_quantity
        puts "The food is cooked"
        switch_to_warm
      else
        puts "Cooking is not ready yet"
      end
    end
  
    def switch_to_warm
      if @cooking
        @cooking = false
        puts "The button has switched to warm"
      else
        puts "Cooking must be finished to switch to warm"
      end
    end
  
    def disconnect_power
      if @power_connected
        @power_connected = false
        puts "The current is no longer flowing"
      else
        puts "The rice cooker is already disconnected"
      end
    end
  
    def exit_program
      puts "Program terminated."
      exit
    end
  end
  
  rice_cooker = RiceCooker.new
  
  begin
    while true
      rice_cooker.display_menu
      print "Choice: "
      choice = gets.to_i
      rice_cooker.handle_user_choice(choice)
    end
  rescue => e
    puts "An error occurred: #{e.message}"
  end
  