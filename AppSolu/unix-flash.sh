tinygo build -target bluepill -o build.bin .
ls /dev/tty.usb* | xargs /Users/paulmaris/Downloads/stm32flash-0.7/stm32flash -w build.bin -v
rm build.bin