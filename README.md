## WakaTime Module for Waybar

### Overview
The <a href='https://wakatime.com/'>WakaTime<a/> Module for Waybar integrates WakaTime with Waybar, allowing you to track your coding activity and display the time spent coding directly on your bar.

> **Note:** To use this module, you must have the WakaTime plugin installed in one of your IDEs. Additionally, a `.wakatime.conf` file should exist in your root directory.

### Installation
1. Download the release archive and unzip it.

2. Grant execute permissions to the binary file:
```bash
  chmod +x ./waybar-code-stats
```
3.Move the binary file to your Waybar configuration directory:
```bash
mv ./waybar-code-stats ~/.config/waybar/
```

4. Configure the module in your Waybar `config.jsonc` file:
```jsonc
"custom/wakatime": {
  "format": "{}{}",
  "return-type": "json",
  // Specify the path to the built executable
  "exec": "~/.config/waybar/waybar-code-stats en", // You can explicitly set the language for display (ru/en)
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
Contributions are welcome! If you encounter any issues or have suggestions for improvements, feel free to:
* Submit an issue on the <a href='https://github.com/AmadoMuerte/WaybarCodeStats'>GitHub repository</a>.
* Open a pull request with your changes.
