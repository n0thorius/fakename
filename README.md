# fakename

It allows you to create simple "fake identity" using [https://fakenamegenerator.com](https://fakenamegenerator.com) and [https://thispersondoesnotexist.com](https://thispersondoesnotexist.com).

Also, it allows you to post this information on the ghostbin, support for pastebin will be added shortly after.

# Installation

```bash
$ go get -u github.com/n0thorius/fakename
```

# Usage

## Generate identity

```bash
$ fakename gen --help
generate fake identity

Usage:
  fakename gen [userCode] [countryCode] [flags]

Flags:
  -d, --download   download image
  -h, --help       help for gen
```

You just need to provide the user set code and the country code. Optionally, you can pass `-d` flag to download some image.

```bash
$ fakename gen us us
Name: Michael H. Kent
Address: 2754 Joes Road, Stephentown, NY 12168

SSN: 061-20-XXXX

Phone: 518-733-0251
Birthday: August 26, 1945
Email: MichaelHKent@rhyta.com
Username: Terew1945
Password: ohkee5aeN

Height: 5' 7" (170 centimeters)
Weight: 149.6 pounds (68.0 kilograms)

```

## Request new image
Just call the command, no need to pass any additional flags. This could be useful in cases where you want other picture not the one previously downloaded.

```bash
$ fakename image
Image '/tmp/fakename/imagecb85ddfb02211006.jpeg' saved successfully
```

## Post to ghostbin (pastebin later)

You can also save this fake identity on ghostbin by passing the name of the file containing information. `gen` command creates the file named as the fake username.

```bash
$ fakename post Terew1945
http://ghostbin.com/paste/a0UWJ
```

## List codes

To see the codes available, pass the `list` command to the binary. You can optinally pass subcommands `u` - user set codes or `c` - country set codes.

```bash
$ fakename list u
			USER SET
Code       | Country
============================
gr         | German
hu         | Hungarian
jp         | Japanese (Anglicized)
no         | Norwegian
fa         | Persian
br         | Brazil
dk         | Danish
fr         | French
sw         | Swedish
ar         | Arabic
er         | Eritrean
fi         | Finnish
hobbit     | Hobbit
th         | Thai
jpja       | Japanese
rucyr      | Russian (Cyrillic)
gd         | Scottish
vn         | Vietnamese
celat      | Chechen (Latin)
en         | England/Wales
sp         | Hispanic
tlh        | Klingon
us         | United States
cs         | Czech
nl         | Dutch
is         | Icelandic
zhtw       | Chinese (Traditional)
ninja      | Ninja
sl         | Slovenian
ch         | Chinese
gl         | Greenland
it         | Italian
pl         | Polish
ru         | Russian
au         | Australian
hr         | Croatian
ig         | Igbo
```

# Note

I know the output is ugly and there are probably more improvements to be done, feel free to let me know any improvements you would like to see.


