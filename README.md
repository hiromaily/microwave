# Microwave
calculate time each wattage for microwave oven

```bash
# microwave [time(s)] [wattage]
microwave 210 500
 wattage: 300 (w), time: 350 (s)
 wattage: 500 (w), time: 210 (s)
 wattage: 600 (w), time: 175 (s)
 wattage: 700 (w), time: 150 (s)
 wattage: 800 (w), time: 131 (s)
 wattage: 1000 (w), time: 105 (s)
 wattage: 1500 (w), time: 70 (s)

# args2:wattage is option, default is 500
microwave 210
 wattage: 300 (w), time: 350 (s)
 wattage: 500 (w), time: 210 (s)
 ...
```

## Install
```
go install github.com/hiromaily/microwave@v1.0.0
```