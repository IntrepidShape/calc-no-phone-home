# A calc that doesn't phone home

A simple calculator application that respects your privacy. It performs calculations locally without sending any data over the network.

## Features

- Basic arithmetic operations (+, -, *, /)
- Exponentiation (^)
- Parentheses support
- Calculation history
- No network communication

## Installation

### Prerequisites

- Go 1.16 or later
- Fyne toolkit

### Steps

1. Install Go from [https://golang.org/](https://golang.org/)

2. Install Fyne:
   ```
   go get fyne.io/fyne/v2
   ```

3. Clone this repository:
   ```
   git clone https://github.com/IntrepidShape/calc-no-phone-home.git
   cd calc-no-phone-home
   ```

4. Build the application:
   ```
   go build
   ```

5. Run the calculator:
   ```
   ./calc-no-phone-home
   ```

## Usage

- Enter expressions using the on-screen buttons or your keyboard
- Press Enter or the "=" button to evaluate
- View calculation history in the scrollable list below
- Click on a history item to load it back into the input field
- Use the "Clear History" button to erase the calculation history

## License

    GLWTS(Good Luck With That Shit) Public License
            Copyright (c) Every-fucking-one, except the Author

Everyone is permitted to copy, distribute, modify, merge, sell, publish,
sublicense or whatever the fuck they want with this software but at their
OWN RISK.

                             Preamble

The author has absolutely no fucking clue what the code in this project
does. It might just fucking work or not, there is no third option.


                GOOD LUCK WITH THAT SHIT PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION, AND MODIFICATION

  0. You just DO WHATEVER THE FUCK YOU WANT TO as long as you NEVER LEAVE
A FUCKING TRACE TO TRACK THE AUTHOR of the original product to blame for
or hold responsible.

IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.

Good luck and Godspeed.