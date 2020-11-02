package load

import "github.com/dhdanie/itinerary-db/model"

var dummyItineraries = []model.Itinerary{
	{
		Id:                  "10",
		StartDate:           "2020-10-09T00:00:00-05:00",
		EndDate:             "2020-10-18T00:00:00-05:00",
		SalutationName:      "Daniel Family",
		SalutationText:      "I’m excited that your trip is almost here!! I wish that I could be your personal guide while there, but since I can’t, I hope you’ll find this document helpful. There is A LOT of information listed below, so please don’t get overwhelmed! You’re going to have an amazing time and make memories that will last forever! I’m thrilled that you get to experience this together! If you need anything, please don’t hesitate to call/text me at 479-313-5901 or send an e-mail to FTJbyCherylDaniel@gmail.com (it will go directly to my phone). I will be having surgery onFriday, December 20 and may be unavailable for a couple of days, but I should be at home recovering when you’re ready to head down to Florida. While you’re at the parks, if you’re unable to reach me for some reason,check with Guest Services, and they will do what they can to make sure your time there is wonderful. If it rains,duck into a shop or an indoor ride, and it will probably be gone within about 15-30 minutes. It will probably be sunny most days (highs in the 70s and lows in the upper 50s to low 60s at night). You can see the forecast here:Orlando weather. Have fun!",
		SalutationSignature: "Cheryl Daniel",
		Infos: []model.Info{
			{
				Title:    "Things to Pack",
				Expanded: false,
				Notes: `
<li>Comfortable shoes/sneakers, shorts , t-shirts, hats, bathing suits (if you plan to swim),pajamas, undergarments
<li>Sweatshirts (it will seem cooler at night)
<li>Deodorant (I forgot this once...no joke...so it gets its own line...haha)
<li>Medicine (prescription and OTC - ibuprofen, Tylenol, Imodium, Bonine (motion sickness forrides), etc.)
<li>Toiletries: toothbrush/toothpaste, hair products/accessories, sanitary items, contacts/solution,Q-tips, soap (resort will have if you don’t need a special kind)
<li>Sunglasses
<li>Camera (if not using cell phone) - check memory card for storage (I left this home on one tripbut I now own a cute Mickey Micro SD card)
<li>Ponchos (in the event of rain - $1 at Walmart or $10+ each at parks)
<li>Sunscreen (SPF 30+ even if it’s cloudy, you can get burned)
<li>Small backpack/bag (whatever is most comfortable) to carry items and always make sure yourbags are zipped, buttoned, etc.
<li>Baby powder or something similar (in case of chafing) and/or something for blisters
<li>Ziplock bags to store wet ponchos and protect phones/electronics on rides - be sure to dry offwet ponchos when you get back to your hotel room
<li>Travel documents/items: IDs, print out documents (if you don’t plan to access them usingphone)
<li>Optional - External battery for cell phone (my battery always runs out at the parks), Wet wipes(small pack - great for quick hand washing before eating), insect repellent (lots of bugs inFlorida although mosquitos almost seem to disappear at the parks), mini first-aid kit
<li>BE SURE TO STAY HYDRATED!!!`,
			},
			{
				Title:    "Things to Remember",
				Expanded: false,
				Notes: `
<li>I recommend NOT doing online check-in as it will override any requests I put into the system.  It also sometimescauses issues with everything working properly.  Sometimes guests who go straight to the parks or DisneySprings will find that things like the dining plan may not work correctly.  SAVE your receipts if you find an errorwith dining and your resort will fix it when you return (this is not something they will allow me to help with onceyou check in).
<li>There are wild animals around Walt Disney World (near water: alligators/snakes/snapping turtles ; on land:snakes, raccoons, armadillos, opossums, bears, etc.).
<li>If you need to get groceries or other essentials, this place can deliver to your resort:(​https://www.gardengrocer.com/​.  My favorite place to shop though is Publix and you should be able to useUber/Lyft/taxi to get there.
<li>You should have received your MagicBands (please let me know if you haven’t).  Inside of the MagicBands box,they now put the ​Magical Extras​ card.  Bring this with you to take advantage of the discounts at stores, mini-golf,etc. There should be something that shares what locations are currently discounted, too.
<li>Keep your bags/purses zipped up and don’t leave ANYTHING valuable out of your site.
<li>Your My Disney Experience APP will have your information listed (make sure to download it if you haven’talready: ​MyDisneyExperience app​)​.  Let me know if you have any questions.
<li>Download the ​Play Disney Parks​ app for some extra fun while there.  You can do interactive things while waitingin lines, earn achievements, and download park music.  Fun for the kids!
<li>When you are at in the Disney parks, you can use the app to continue to choose Fast Passes after your initialthree are used up and there are others still available (check right after you scan your third one while waiting inline - you can book it then).  Continue to “refresh” for better times or select one and then “modify” it.  It may giveyou times closer to the time you are currently at (switch the time to the current time to see if anything isavailable).  You can also use the app to check balance of dining credits remaining, the wait time for buses, orderquick service meals before arriving, and more.
<li>With your dining plan, you have 6 table service dining credits, 6 quick service dining credits, and 12 snackcredits (look for the purple icon on menus) per person to use per night’s stay (these can be used on check-in andcheck-out days, too).   Just scan your band to use the credits.  As mentioned above, save your receipts in casethere is an error and you’re charged when you should not be. You will get your refillable beverage mug at yourresort (usually at the quick service location). They will scan your band to redeem it. A quick service credit can beturned into 3 snack credits (look for those purple icons) IF it’s done in one transaction at a given location (forexample, if instead of using a quick service credit for breakfast, you want a muffin, a banana, and a milk, you cando that).
<li>If you want to cancel a dining reservation, in most cases, they need 24 hours notice.  If not given, they willcharge $10 per person for the reservation held with your credit card.
<li>You get more bang for your buck if you add an alcoholic beverage or a specialty beverage (milkshake, etc.) toyour meal.  Then, you can get water or soda and choose to pay for that.  You’ll only receive one alcoholicbeverage/specialty drink with your meal. If you’re eating at the resort and using your mug with a quick servicemeal, grab a bottled water or bottle of soda (it comes with the meal anyway).
<li>You did not purchase the MemoryMaker.  If you decide you’d like to add this, it’s $169 in advance or $199 onceat the park.  The MemoryMaker will include photos taken on some of the rides and from park photographers(digital downloads). If you decide to add this, be sure to ask for “MAGIC” shots (if they don’t suggest them).   Youcan start viewing pictures in your My Disney Experience app soon after the first pictures are taken (can take up to24 hours to show up), and once you start viewing the pictures, you will have 45 days to download them, so makesure to download them when back home.  To ACTIVATE the MemoryMaker in the app, click on any photo andpush the activate button.  Then, you can start viewing pictures without the watermark.
<li>If at any point you decide you don’t want to wait on the buses, there is a Minnie Van service (Lyft) where youcan pay extra for this option.  Here is a link with more info:https://disneyworld.disney.go.com/en-eu/guest-services/minnie-van-service/​.`,
			},
		},
		Attractions: []model.Attraction{
			{
				Date:    "2020-10-10T09:01:00-05:00",
				VenueId: 14,
				Notes:   "Get ready for the MAGIC at the most MAGICAL PLACE ON EARTH!!  Once you enter the Magic Kingdom, you'll head straight down Main Street where you'll smell tasty treats and window shop for souvenirs galore!  As you approach the castle, you will notice that it branches off in different directions at the hub near the Partners (Walt & Mickey) statue.<br><br> To the left, you will see Adventureland where you can visit classic attractions like Pirates of the Caribbean and Jungle Cruise.  Check out Aloha Isle if you think you might enjoy one of the most popular Disney treats, a Dole Whip!  Young children may also enjoy soaring high while flying on the Magic Carpet ride.<br><br> Between Adventureland and Cinderella Castle, another path will take you towards Frontierland where you can try popular rides like Splash Mountain (closing soon to be reimagined as The Princess and the Frog theme) and Big Thunder Mountain Railroad.  If you're feeling adventurous, it's just a hop, skip, and a jump across the water to visit Tom Sawyer's Island.<br><br> Also in that direction, you'll find Liberty Square where the Haunted Mansion and Hall of Presidents is found (we love to watch the show while we cool off from the heat or hide out until the rain passes).  Sleepy Hollow will have treats that fill you up like a meal, and they often has seasonal<br><br> Straight through the castle, you'll find Fantasyland which is filled with attractions like Peter Pan, it's a small world, 7 Dwarfs Mine Train (one of the most popular roller coasters), and so much more.  You'll also find some areas where you may see some characters out and about, so be sure to look around, and wave hello if you see someone!  If you're looking for a fun show to cool off inside here, check out Mickey's PhilharMagic where you'll enjoy some familiar tunes and special effects that will immerse you in the experience.<br><br> Finally, if you're facing the castle, from Main Street, towards the right there's a path that will lead To Infinity and Beyond, okay, it's actually tomorrowland, but you'll be able to check out Buzz Lightyear's Space Ranger Spin, Space Mountain, and more.<br><br> This theme park is always TOP on my list of places to visit, and I sure do hope you'll love it, too!",
			},
		},
		HotelStays: []model.HotelStay{
			{
				CheckInDate:    "2020-10-09T21:25:00-05:00",
				CheckOutDate:   "2020-10-17T16:58:00-05:00",
				ConfirmationNo: "2078910320789103",
				VenueId:        13,
				Notes: `
<li>About the Area - Located in Lake Buena Vista's Bay Lake neighborhood, Disney's Grand Floridian Resort & Spa is near the airport and on a lake. Grand Cypress Golf Club and Disney's Lake Buena Vista Golf Course are worth checking out if an activity is on the agenda, while those in the mood for shopping can visit Disney Springs® and Old Town. Traveling with kids? Don't miss Walt Disney World®, or check out an event or a game at ESPN Wide World of Sports. Our guests love the resort's quiet location.
<li>Property Location - This family-friendly Lake Buena Vista resort is located near the airport, within 1 mi (2 km) of Disney's Palm Golf Course and Disney's Magnolia Golf Course. Magic Kingdom® Park and Epcot® are also within 6 mi (10 km).
<li>Rooms - Disney's Grand Floridian Resort & Spa's 867 rooms are air-conditioned and provide balconies, MP3 docks, and coffee makers. Guests can expect to find free WiFi and TVs with cable channels. Ceiling fans, hair dryers, and safes are also standard.
<li>Amenities - Guests of Disney's Grand Floridian Resort & Spa have access to a full-service spa, 2 outdoor pools, and a children's pool. Valet and self parking are available, and there's also limo/town car service. Multilingual staff at the front desk are standing by 24/7 to help with concierge services, dry cleaning/laundry, and luggage storage. Other amenities at this beach resort include a fitness center, an arcade/game room, and express check-out.
<li>Dining - Enjoy dining at Victoria and Albert, one of 7 onsite restaurants. Order from room service to satisfy your hunger without leaving your room, or explore the resort and get a beverage at the bar/lounge or coffee shop/café. Start each morning with cooked-to-order breakfast, available for a fee from 6:30 AM to 11:30 AM.
<li>Business, Other Amenities - Featured amenities include limo/town car service, express check-out, and dry cleaning/laundry services. Self parking (subject to charges) is available onsite.`,
			},
		},
		DiningReservations: []model.DiningReservation{
			{
				Date:           "2020-10-10T13:35:00-05:00",
				ConfirmationNo: "502251810044",
				VenueId:        15,
				Notes:          "Enter the Beast’s Enchanted Castle for a quick-service breakfast, lunch or an unforgettable sit-down dinner—bon appétit!<br><br>If permitted, ask to sit in the West Wing if you think it won't be too scary.  It's a bit darker in here, but you'll be able to watch the rose and see changes in the paintings when you hear the rumble of thunder and the flickering of the lights when lightning flashes.",
			},
		},
		Flights: []model.Flight{
			{
				Airline:          "American Airlines",
				FlightNumber:     "343",
				DepartureDate:    "2020-10-09T14:15:00-05:00",
				DepartureVenueId: 10,
				ArrivalDate:      "2020-10-09T16:58:00-05:00",
				ArrivalVenueId:   11,
				Terminal:         "C3",
				RecordLocator:    "YK7G44",
			},
			{
				Airline:          "American Airlines",
				FlightNumber:     "DL1555",
				DepartureDate:    "2020-10-09T19:57:00-05:00",
				DepartureVenueId: 11,
				ArrivalDate:      "2020-10-09T21:25:00-05:00",
				ArrivalVenueId:   12,
				BaggageClaim:     "A2",
				RecordLocator:    "YK7G44",
			},
		},
	},
}
