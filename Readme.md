# Amazon Scrapper

A blazing fast amazon scrapper using golang and colly.

Returns JSON data.

## How to execute it

1. Make sure you have [golang](https://golang.org/doc/install) installed.

2. Open the terminal in project folder and type: `go get -u github.com/gocolly/colly/...`

3. Run the script with `go run main.go`

### Example output:

```
[
  {
    "ProductName": "2020 Nintendo Switch with Neon Blue and Red Controllers w/ 69 Value HESVAP 13in1 Supper Kit Case (Earphone, LCD Film, Card Case, Silicon Case x 2pcs, Carry Bag, Wiping Cloth etc.)",
    "Stars": "4.4",
    "Price": "549.99"
  },
  {
    "ProductName": "Nintendo Switch Lite - Coral - Switch",
    "Stars": "4.8",
    "Price": "314.96"
  },
  {
    "ProductName": "Nintendo Switch w/ Gray Joy-Con + Mario Kart 8 Deluxe (Full Game Download) - Switch",
    "Stars": "4.7",
    "Price": "473.99"
  },
  {
    "ProductName": "Joy Con Controller Replacement Campatiable for Nintendo Switch - Left and Right Controllers Compatible for Nintendo Switch Console, Wired/Wireless L/R Switch Remotes",
    "Stars": "3.8",
    "Price": "38.99"
  },
  {
    "ProductName": "Locking Carry Case for Nintendo Switch Hardshell Deluxe Bag w/Anti-Theft TSA Combination Lock - Include 1 Tempered Glass Screen Protector",
    "Stars": "4.8",
    "Price": "32.99"
  },
  {
    "ProductName": "Nintendo Switch Lite - Turquoise",
    "Stars": "4.7",
    "Price": "270.00"
  },
  {
    "ProductName": "Pikmin 3 Deluxe - Nintendo Switch",
    "Stars": "Unknown",
    "Price": "59.99"
  },
  {
    "ProductName": "Nintendo 32GB Switch with Gray Joy-Con Controllers - with SanDisk 128GB UHS-I microSDXC Memory Card for The Switch",
    "Stars": "4.5",
    "Price": "425.00"
  },
  {
    "ProductName": "2020 Nintendo Switch with Neon Blue and Red Controllers w/ 69 Value HESVAP 13in1 Supper Kit Case (Earphone, LCD Film, Card Case, Silicon Case x 2pcs, Carry Bag, Wiping Cloth etc.)",
    "Stars": "4.4",
    "Price": "549.99"
  },
  {
    "ProductName": "Paper Mario: The Origami King - Nintendo Switch",
    "Stars": "4.6",
    "Price": "58.00"
  },
  {
    "ProductName": "Ring Fit Adventure - Nintendo Switch",
    "Stars": "4.8",
    "Price": "79.99"
  },
  {
    "ProductName": "Mario Kart 8 Deluxe - Nintendo Switch",
    "Stars": "4.9",
    "Price": "49.94"
  },
  {
    "ProductName": "Charger for Nintendo Switch，Nintendo Switch AC Adapter,Hdmi Splitter 1 in 1 Out,USB3.0，3in1 USB Quick Charger (Support TV Mode)",
    "Stars": "5.0",
    "Price": "39.99"
  },
  {
    "ProductName": "Switch Accessories Bundle - Orzly Geek Pack for Nintendo Switch: Case \u0026 Screen Protector, Joycon Grips \u0026 Racing Wheels, Switch Controller Charge Dock, Comfort Grip Case \u0026 More - JetBlack",
    "Stars": "4.7",
    "Price": "54.99"
  },
  {
    "ProductName": "Switch Bluetooth Adapter for Nintendo Switch/Switch Lite/PS4/PC, Dual Stream Bluetooth Wireless USB C Audio Transmitter with aptX Low Latency to AirPods Bluetooth Speakers Headphones",
    "Stars": "5.0",
    "Price": "27.99"
  }
]
```