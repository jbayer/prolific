package main

import "io/ioutil"

const template = `As a user I can toast a bagel

When I insert a bagel into toaster and press the on button, I should get a toasted bagel

L: mvp, toasting

---

As a user I can set the desired color of my bagel

I should be able to manipulate a dial and choose one of:

- light
- medium
- dark

Pressing the on button gives me toast of the appropriate color.

L: mvp, toasting

---

As a user I can clean my bagel toaster

I should be able to pull out a tray and clean up the crumbs.

L: mvp, clean-up`

func GenerateTemplate() error {
	return ioutil.WriteFile("stories.prolific", []byte(template), 0666)
}
