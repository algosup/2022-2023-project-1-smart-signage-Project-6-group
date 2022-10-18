 <center> <h1> AppSolu </h1> </center>

<strong>Prepared by :</strong>
<br><strong>Lucas Aubard</strong></br>

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Turn on the board | 1 | 1.Plug the board | The board turn on | the electric current passes through the board then the board turns on | no problem to turn on the board |

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Turn off the board | 2 | 1.Go to the terminal 2.Write"xxxx"+"off" | The board turn on |  |  |
|2|Turn off the board|2|1.Go to the terminal 2.Write"xxxx"+"number"|The board don't crash|||
|3|Turn off the board|2|1.Go to the terminal 2.Write"xxxx"+"string"|The board don't crash|||
|4|Turn off the board|2|1.Go to the terminal 2.Write"xxxx"+"_"|The board don't crash|||

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Turn the led on | 2 | 1.Go to the terminal 2.Write"xxxx"+"on" | The led turn on |  |  |
|2|Turn the led on|2|1.Go to the terminal 2.Write"xxxx"+"number"|The board don't crash|||
|3|Turn the led on|2|1.Go to the terminal 2.Write"xxxx"+"string"|The board don't crash|||
|4|Turn the led on|2|1.Go to the terminal 2.Write"xxxx"+"_"|The board don't crash|||

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Turn the led off | 2 | 1.Go to the terminal 2.Write"xxxx"+"off" | The led turn off |  |  |
|2|Turn the led off|2|1.Go to the terminal 2.Write"xxxx"+"number"|The board don't crash|||
|3|Turn the led off|2|1.Go to the terminal 2.Write"xxxx"+"string"|The board don't crash|||
|4|Turn the led off|2|1.Go to the terminal 2.Write"xxxx"+"_"|The board don't crash|||

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Change led's brightness | 2 | 1.Go to the terminal 2.Write"xxxx"+"100%" | the brightness is 100% |  |  |
|2|Change led's brightness|2|1.Go to the terminal 2.Write"xxxx"+"60%"|the brightness is 60%|||
|3|Change led's brightness|2|1.Go to the terminal 2.Write"xxxx"+"40%"|the brightness is 40%|||
|4|Change led's brightness|2|1.Go to the terminal 2.Write"xxxx"+"0%"|the brightness is 0%|||
|5|Change led's brightness|2|1.Go to the terminal 2.Write"xxxx"+"-100%"|the brightness is 0% and the board don't crash|||
| 6 | Change led's brightness | 2 | 1.Go to the terminal 2.Write"xxxx"+"string" | the brightness don't change and the board don't crash |  |  |
| 7 | Change led's brightness | 2 | 1.Go to the terminal 2.Write"xxxx"+"_" | the brightness don't change and the board don't crash |  |  |

| N° | ACTION            | NUMBER OF STEPS | STEPS DESCRIPTION | EXPECTATION | RESULT | WHAT’S THE PROBLEM |
| -- | ----------------- | --------------- | ----------------- | ----------- | ------ | ------------------ |
| 1 | Check Battery level | 2 | 1.Go to the terminal 2.Write"xxxx"+"return battery" | Return battery level |  |  |
|2|Check Battery level|2|1.Go to the terminal 2.Write"xxxx"+"number"|The board don't crash|||
|3|Check Battery level|2|1.Go to the terminal 2.Write"xxxx"+"string"|The board don't crash|||
|4|Check Battery levelTurn on the board|2|1.Go to the terminal 2.Write"xxxx"+"_"|The board don't crash|||
