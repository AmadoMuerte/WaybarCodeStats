## WakaTime Module for Waybar

### Overview
The WakaTime Module for Waybar integrates WakaTime with Waybar, allowing you to track your coding activity and display the time spent coding directly on your bar.

> **Note:** To use this module, you must have the WakaTime plugin installed in one of your IDEs. Additionally, a `.wakatime.conf` file should exist in your root directory.

### Installation
1. Clone this repository into your Waybar configuration directory (typically `~/.config/waybar/`):

```bash
# You can build a standalone executable file using Go, if desired
# Please ensure you have Go installed first
cd WakaTimeModule && go mod tidy && go build ./cmd/main.go
# This will generate an executable file named 'main' in the root of the module
```

2. Configure the module in your Waybar `config.jsonc` file:

```jsonc
"custom/wakatime": {
  "format": "{}{}",
  "return-type": "json",
  // Specify the path to the built executable
  "exec": "~/.config/waybar/WakaTimeModule/main en", // You can explicitly set the language for display (ru/en)
  // Alternatively, run the Go file directly
  // "exec": "cd ~/.config/waybar/WakaTimeModule && go run ./cmd/main.go en"
  "interval": 1200
},
```

### Customization
Once installed, you can customize the module's appearance in your `style.css` file:
```css
#custom-wakatime {
color: #a46df6; /* Adjust the color as desired */
border-radius: 20px;
margin-left: 5px;
border-right: 0px;
transition: 0.3s;
padding-left: 7px;
}
```

### Contribution
Feel free to submit issues or pull requests to contribute to the development of this module.
