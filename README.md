# **SIMPOS BOT**

SimPOS Bot, or spb for short, is a CLI testing automation tool. 
It utilizes the SimPOS REST API to automate testing efforts.
Its main strengths are light-weight and extensible.
## **Usage**

    spb run --file <config file>
Run predefined test cases in a config file.

    spb run --queue reversal/adjustment/both
Run either the reversal, adjustment queue, or both.

## **Downloads**

https://github.com/anhdvu/simposbot/releases

## **Config File**

A config file is typically a .yaml file. [Learn more about yaml file](https://en.wikipedia.org/wiki/YAML).

However, any file extension should work as long as the content follows YAML syntax.

### **Sample**

    name: Sample config
    cookie:
    - CFID=20557
    - CFTOKEN=3b79a1e3c773c94-E7F38875-D295-8B7C-22E05E8FB40599AD
    testcard:
      number: '5338485258218895'
      expirydate: '3011'
      cvv: '654'
      pin: '1234'
    shared:
      amountMin: 1
      amountMax: 6
      defaultOriginalCurrencyCode: '978'
      defaultOriginalCurrencyDecimalPlaces: '2'
      defaultProvince:
      defaultCountry:
      token:
    testcases:
    - included: false 
      name: pos reversal
      runs: 1 
      mode: pos 
      function: 
      atm: false  
      source: nfc 
      foreign: false 
      originalCurrencyCode: 
      originalCurrencyDecimalPlaces: 
      acquirer: POS DEDUCT REV
      province: COMPANION
      country: DEU
      mcc: '7299'
      reversal: full
      advice: false  


***IMPORTANT***: *The hyphen "-" before "included" is mandatory because it signals a new test case.*

### **Test Case Settings**

- **included**: 
> Possible values: *true*/*false*. Set to *false* to exclude the test in the run. Default value: *true*.
- **name**: 
> Name of the test case.
- **runs**: 
> Set the number of runs for a specific test case. Default value: 1.
- **mode**: 
> Possible values: *pos*, *web*, *settlement*, *payment*. Any other values will throw error.
- **function**:
> This field is tied with field mode above, therefore it must be set accordingly. Otherwise, it will throw error. 
> Possible values:
> - *pos*: *deduct*
> - *web*: *deduct*
> - *settlement*: *refund*, *fxload*, *fxdeduct*, *noauth*, *chargeback*
> - *payment*: *payment*, *refund*
- **atm**: 
> Possible values: *true/false*. Set to *true* to indicate the transaction type is ATM (01). Default value: *false*.
- **source**: 
> Possible values: *mag*, *emv*, or *nfc*. Any other values will be assumed as *emv*. Default value: *emv*.
- **foreign**: 
> Possible values: *true*/*false*. Set to *true* to indicate the original currency is not the one set in campaign setting. Default value: *false*.
- **originalCurrencyCode**: 
> If not explicitly set while foreign == *true*, the tool will take defaultOriginalCurrencyCode instead.
- **originalCurrencyDecimalPlaces**: 
> If not explicitly set while foreign == *true*, the tool will take defaultOriginalCurrencyDecimalPlaces instead.
- **acquirer**:
> Any string should work. If it is longer than 22 characters, it will be automatically truncated.
- **province**:
> Any string should work. If it is longer than 13 characters, it will be automatically truncated.
- **country**: 
> Any string should work. It will truncated if it's longer than 3 characters. Note: Country should be set to align with originalCurrencyCode if foreign == *true*.
- **mcc**:
> The tool does NOT check for valid MCC. Any 4 digits should work
- **reversal**: 
> Possible values: *partial* or *full*. Any other values will be assumed as "". Default value: "" (empty string).
> Note: if paymentType is set, *partial* and *full* will have the same behavior which is full reversal.
- **advice**: 
> Possible values: *true*/*false*. Set to *true* to indicate whether a transaction is a Deduct Advice. This field is irrelevant when mode == *Settlement*. Default value: *false*.

### **Other Settings**

- **name**: Any name is fine.

- **cookie**: This is used to retrive JWT, which, in turn, is used to run test cases with SimPOS. You can set your own values by getting them via your browser's DevTools.
- **testcard**: card details are set here.
  - number
  - expirydate
  - cvv
  - pin

- **shared**: configuration that is shared among test cases.
  > The tool will automatically generate an amount within given bounds of 2 values below.
  - amountMin: 
  - amountMax:
  > The tool will use default values below if they are not explicitly set in test case settings.
  - defaultOriginalCurrencyCode:
  - defaultOriginalCurrencyDecimalPlaces: