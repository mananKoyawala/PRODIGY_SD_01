# Temperature Converter

## Description

This is the first task of my internship at Prodigy Infotech. The goal of this task is to create a program that converts temperatures between Celsius, Fahrenheit, and Kelvin scales. The program prompts the user to input a temperature value and the original unit of measurement. It then converts the temperature to the other two units and displays the converted values to the user.

## Features

- Convert temperature from Celsius to Fahrenheit and Kelvin.
- Convert temperature from Fahrenheit to Celsius and Kelvin.
- Convert temperature from Kelvin to Celsius and Fahrenheit.
- User-friendly command-line interface with color-coded output.
- Input validation to ensure proper numerical input.

## Installation

1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/temperature-converter.git
   cd temperature-converter
   ```
2. **Install dependencies:**
   This application uses the github.com/fatih/color package for colored output. Install it using:
   ```
   go get github.com/fatih/color
   ```
3. **Run the application:**
   ```
   go run main.go
   ```

## Usage

1. Run the application.
2. You will be presented with a menu to choose the temperature unit to convert from:
   1: Celsius
   2: Fahrenheit
   3: Kelvin
   4: Exit the application
3. Enter the temperature value you want to convert.
4. The application will display the converted values in the other two units.

## Example

If you enter a temperature of 25 degrees Celsius, the application will convert it to Fahrenheit and Kelvin, and present the converted values as outputs.

```
----- Temperature Converter -----
1. Enter temperature in Celsius
2. Enter temperature in Fahrenheit
3. Enter temperature in Kelvin
4. Exit from app

Enter your choice
1

Enter Celsius
25

Input Celsius value is 25.00
Fahrenheit is 77.00 and Kelvin is 298.15
```

## Contributing

Feel free to fork this project and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT License](LICENSE)
