package avatar

import (
	"fmt"
	"math/rand"
	"time"
)

func Url() string {
	url := fmt.Sprintf(
		"https://avataaars.io/?clotheColor=%s&accessoriesType=%s&avatarStyle=%s&clotheType=%s&eyeType=%s&eyebrowType=%s&facialHairColor=%s&facialHairType=%s&hairColor=%s&hatColor=%s&mouthType=%s&skinColor=%s&topType=%s",
		clotheColor(),
		accessoriesType(),
		avatarStyle(),
		clotheType(),
		eyeType(),
		eyebrowType(),
		facialHairColor(),
		facialHairType(),
		hairColor(),
		hatColor(),
		mouthType(),
		skinColor(),
		topType())

	return url
}

func clotheColor() string {
	random := make(map[int]string, 0)
	random[0] = "Black"
	random[1] = "Blue01"
	random[2] = "Blue02"
	random[3] = "Blue03"
	random[4] = "Gray01"
	random[5] = "Gray02"
	random[6] = "Heather"
	random[7] = "PastelBlue"
	random[8] = "PastelGreen"
	random[9] = "PastelOrange"
	random[10] = "PastelRed"
	random[11] = "PastelYellow"
	random[12] = "Pink"
	random[13] = "Red"
	random[14] = "White"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(15)]
}

func accessoriesType() string {
	random := make(map[int]string, 0)
	random[0] = "Blank"
	random[1] = "Kurt"
	random[2] = "Prescription01"
	random[3] = "Prescription02"
	random[4] = "Round"
	random[5] = "Sunglasses"
	random[6] = "Wayfarers"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(7)]
}

func avatarStyle() string {
	as := make(map[int]string, 0)
	as[0] = "Circle"
	as[1] = "Transparent"
	return as[rand.Intn(2)]
}

func clotheType() string {
	random := make(map[int]string, 0)
	random[0] = "BlazerShirt"
	random[1] = "BlazerSweater"
	random[2] = "CollarSweater"
	random[3] = "GraphicShirt"
	random[4] = "Hoodie"
	random[5] = "Overall"
	random[6] = "ShirtCrewNeck"
	random[7] = "ShirtScoopNeck"
	random[8] = "ShirtVNeck"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(9)]
}

func eyeType() string {
	random := make(map[int]string, 0)
	random[0] = "Close"
	random[1] = "Cry"
	random[2] = "Default"
	random[3] = "Dizzy"
	random[4] = "EyeRoll"
	random[5] = "Happy"
	random[6] = "Hearts"
	random[7] = "Side"
	random[8] = "Squint"
	random[9] = "Surprised"
	random[10] = "Wink"
	random[11] = "WinkWacky"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(12)]
}

func eyebrowType() string {
	random := make(map[int]string, 0)
	random[0] = "Angry"
	random[1] = "AngryNatural"
	random[2] = "Default"
	random[3] = "DefaultNatural"
	random[4] = "FlatNatural"
	random[5] = "RaisedExcited"
	random[6] = "RaisedExcitedNatural"
	random[7] = "SadConcerned"
	random[8] = "SadConcernedNatural"
	random[9] = "UnibrowNatural"
	random[10] = "UpDown"
	random[11] = "UpDownNatural"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(12)]
}

func facialHairColor() string {
	random := make(map[int]string, 0)
	random[0] = "Auburn"
	random[1] = "Black"
	random[2] = "Blonde"
	random[3] = "BlondeGolden"
	random[4] = "Brown"
	random[5] = "BrownDark"
	random[6] = "Platinum"
	random[7] = "Red"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(8)]
}

func facialHairType() string {
	random := make(map[int]string, 0)
	random[0] = "Blank"
	random[1] = "BeardMedium"
	random[2] = "BeardLight"
	random[3] = "BeardMajestic"
	random[4] = "MoustacheFancy"
	random[5] = "MoustacheMagnum"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(6)]
}

func hairColor() string {
	random := make(map[int]string, 0)
	random[0] = "Auburn"
	random[1] = "Black"
	random[2] = "Blonde"
	random[3] = "BlondeGolden"
	random[4] = "Brown"
	random[5] = "BrownDark"
	random[6] = "PastelPink"
	random[7] = "Blue"
	random[8] = "Platinum"
	random[9] = "Red"
	random[10] = "SilverGray"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(11)]
}

func hatColor() string {
	random := make(map[int]string, 0)
	random[0] = "Black"
	random[1] = "Blue01"
	random[2] = "Blue02"
	random[3] = "Blue03"
	random[4] = "Gray01"
	random[5] = "Gray02"
	random[6] = "Heather"
	random[7] = "PastelBlue"
	random[8] = "PastelGreen"
	random[9] = "PastelOrange"
	random[10] = "PastelRed"
	random[11] = "PastelYellow"
	random[12] = "Pink"
	random[13] = "Red"
	random[14] = "White"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(15)]
}

func mouthType() string {
	random := make(map[int]string, 0)
	random[0] = "Concerned"
	random[1] = "Default"
	random[2] = "Disbelief"
	random[3] = "Eating"
	random[4] = "Grimace"
	random[5] = "Sad"
	random[6] = "ScreamOpen"
	random[7] = "Serious"
	random[8] = "Smile"
	random[9] = "Tongue"
	random[10] = "Twinkle"
	random[11] = "Vomit"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(12)]
}

func skinColor() string {
	random := make(map[int]string, 0)
	random[0] = "Tanned"
	random[1] = "Yellow"
	random[2] = "Pale"
	random[3] = "Light"
	random[4] = "Brown"
	random[5] = "DarkBrown"
	random[6] = "Black"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(7)]
}

func topType() string {
	random := make(map[int]string, 0)
	random[0] = "NoHair"
	random[1] = "Eyepatch"
	random[2] = "Hat"
	random[3] = "Hijab"
	random[4] = "Turban"
	random[5] = "WinterHat1"
	random[6] = "WinterHat2"
	random[7] = "WinterHat3"
	random[8] = "WinterHat4"
	random[9] = "LongHairBigHair"
	random[10] = "LongHairBob"
	random[11] = "LongHairBun"
	random[12] = "LongHairCurly"
	random[13] = "LongHairCurvy"
	random[14] = "LongHairDreads"
	random[15] = "LongHairFrida"
	random[16] = "LongHairFro"
	random[17] = "LongHairFroBand"
	random[18] = "LongHairNotTooLong"
	random[19] = "LongHairShavedSides"
	random[20] = "LongHairMiaWallace"
	random[21] = "LongHairStraight"
	random[22] = "LongHairStraight2"
	random[23] = "LongHairStraightStrand"
	random[24] = "ShortHairDreads01"
	random[25] = "ShortHairDreads02"
	random[26] = "ShortHairFrizzle"
	random[27] = "ShortHairShaggyMullet"
	random[28] = "ShortHairShortCurly"
	random[29] = "ShortHairShortFlat"
	random[30] = "ShortHairShortRound"
	random[31] = "ShortHairShortWaved"
	random[32] = "ShortHairSides"
	random[33] = "ShortHairTheCaesar"
	random[34] = "ShortHairTheCaesarSidePart"
	rand.Seed(time.Now().UnixNano())
	return random[rand.Intn(35)]
}
