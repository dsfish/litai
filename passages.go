package main

type passage struct {
	author string
	year int
	title string
	text string
}

var passages = []passage{
	{
		author: "Gottfried Leibniz",
		year: 1666,
		title: "On the Combinatorial Art",
		text: "The art of combinations in particular, " +
			"as I take it (it can also be called a general characteristic or algebra), " +
			"is that science in which are treated the forms or formulas of things in general, that is, " +
			"quality in general or similarity and dissimilarity; in the same way that ever new formulas arise from" +
			" the elements a, b, c themselves when combined with each other, " +
			"whether these elements represent quantities or something else.",
	},
	{
		author: "Thomas Dunn English",
		year: 1890,
		title: "Skeleton Essays",
		text: " Hence the preparation of this book, which will supply a need long existing, " +
			"and a demand which though moderate, is steady, and sure to increase. It presents a number of skeletons," +
			" or outlines, of essays on a variety of topics, " +
			"and in some instances these are followed by completed work, in order to show the reader how it is done.",
	},
	{
		author: "Gustav Freytag",
		year: 1900,
		title: "Technique of the Drama",
		text: "The poet of the present is inclined to look with amazement upon a method of work in which the" +
			" structure of scenes, " +
			"the treatment of characters and the sequence of effects were governed by a transmitted code of fixed" +
			" technical rules. Such a limitation easily seems to us the death of free artistic creation.",
	},
	{
		author: "Gottfried Leibniz",
		year: 1679,
		title: "Philosophical Papers and Letters",
		text: "Now, however, our characteristic will reduce the whole to numbers, " +
			"so that reasons can also be weighted as if by a kind of statics for probabilities too, " +
			"will be treated in this calculation and demonstration, " +
			"since one can always estimate which of the given circumstances will more probably occur. Finally, " +
			"anyone who is certainly convinced of the truth of religion and its consequences, " +
			"and so embraces others in love that he desires this conversion of mankind. will surely admit, " +
			"if he understands this matters, " +
			"that nothing will be more influential than this discovery for the propagation of the faith.",
	},
	{
		author: "Gustav Freytag",
		year: 1900,
		title: "Technique of the Drama",
		text: "If, however, we must deny ourselves the advantage of composing according to the craftsman's" +
			" traditions which were peculiar to the dramatic art as well as the plastic art of former centuries, " +
			"yet we should not scorn to seek, and intelligently to use, " +
			"the technical rules of ancient and modern times, which facilitate artistic effects of our stage.",
	},
	{
		author: "Aristotle",
		year: -335,
		title: "Poetics",
		text: "We therefore have (i) the mimesis of the action, the plot, " +
			"by which I mean the ordering of the particular actions; (" +
			"i) the mimesis of the moral characters of the personages, " +
			"namely that in the play which makes us say that the agents have certain moral qualities; (" +
			"ill) the mimesis of their intellect, " +
			"namely those parts in which they demonstrate something in speech or deliver themselves of some general" +
			" maxim.",
	},
	{
		author: "Quirinus Kuhlmann",
		year: 1671,
		title: "Love-Kiss XLI",
		text: "From Night / Fog / Clash / Frost / Wind / Sea / Heat / South / East / West / North/ Sun." +
			" Fire and Plagues / Come Day / Blaze / Bloom / Snow/ Peace / Land / Flash / Warmth / Heat / Joy/ Cool" +
			" / Light/ Flames and Dread/",
	},
	{
		author: "Georges Polti",
		year: 1916,
		title: "The Thirty-Six Dramatic Situations",
		text: "Which are the dramatic situations neglected by our own epic, " +
			"so faithful in repeating the few most familiar? Which, on the other hand, " +
			"are most in use today? Which are the most neglected, and which the most used, in each epoch, genre, " +
			"school, author? What are the reasons for these preferences? The same questions may be asked before the" +
			" classes and subclasses of the situations.",
	},
	{
		author: "Georges Polti",
		year: 1916,
		title: "The Thirty-Six Dramatic Situations",
		text: "It is equally natural that only the greatest and most complete civilizations should have evolved" +
			" their own particular conception of the drama, " +
			"and that one of these new conceptions should be revealed by each new evolution of society, " +
			"whence arises the dim but faithful expectation of our own age, " +
			"waiting for the manifestation of its own dramatic ideals, " +
			"before the cenotaphs of an art which has long been, apparently for commercial reasons, " +
			"almost non-existent.",
	},
	// {
	// 	author: "John Wilkins",
	// 	year: 1668,
	// 	title: "An Essay Towards a Real Character, and a Philosophical Language",
	// 	text: "The second Part shall contain that which is the great foundation of the thing here designed," +
	// 		" namely a regular enumeration and description of all those things and notions, " +
	// 		"to which marks or names ought to be assigned according to their respective natures, " +
	// 		"which may be styled the Scientific Part, comprehending Universal Philosophy. " +
	// 		"It being the proper end and design of the several branches of Philosophy to reduce all things and" +
	// 		" notions unto such a frame, as may express their natural order, dependence, " +
	// 		"and relations [..] The design of this Treatise being an attempt towards a new kind of Character and" +
	// 		" Language, it cannot therefore be improper to premise somewhat concerning those already in being; the" +
	// 		" first Original of them, their several kinds, " +
	// 		"the various changes and corruptions to which they are liable, " +
	// 		"together with the manifold defects belonging to them.",
	// },
	{
		author: "Ramon Llull",
		year: 1308,
		title: "Ars Brevis",
		text: "And thus the intellect has a ladder for ascending and descending; as, for instance, " +
			"descending from a completely general principle to one neither completely general nor completely" +
			" particular, and from a principle neither completely general nor completely particular to one that is" +
			" completely particular. And in a similar fashion one can discuss the ascent of this ladder ." +
			"] Everything that exists is implicit in the principles of this figure, " +
			"for everything is either good or great, etc., as God and angels, which are good, great, etc. Therefore," +
			" whatever exists is reducible to the above-mentioned principles.",
	},
	{
		author: "Ramon Llull",
		year: 1308,
		title: "Ars Brevis",
		text: "We have employed an alphabet in this Art so that it can be used to make figures, " +
			"as well as to mix principles and rules for the purpose of investigating the truth. For, " +
			"as a result of any one letter having many meanings, " +
			"the intellect becomes more general in its reception of the things signified, " +
			"as well as in acquiring knowledge. And this alphabet must be learned by heart, " +
			"for otherwise the artist will not be able to make proper use of this Art.",
	},
	{
		author: "Quirinus Kuhlmann",
		year: 1671,
		title: "Love-Kiss XLI",
		text: "At first glance, what is printed here above/seems impossible; yet just as certainly as two times two" +
			" makes six / so this too is the case. It might seem even more implausible/and here- I have the greatest" +
			" authorities to back me up / and can only argue this briefly here/though at greater length elsewhere" +
			"/ that the major portion of Human Knowledge in fact lies hidden in permutation.",
	},
	{
		author: "Samuel Butler",
		year: 1872,
		title: "Erewhon",
		text: "Even a potato in a dark cellar has a certain low cunning about him which serves him in excellent" +
			" stead. He knows perfectly well what he wants and how to get it. " +
			"He sees the light coming from the cellar window and sends his shoots crawling straight thereto: they" +
			" will crawl along the floor and up the wall and out at the cellar window, " +
			"if there be a little earth anywhere on the journey he will find it and use it for his own ends.",
	},
	{
		author: "Athanasius Kircher",
		year: 1674,
		title: "Letter", // from Gerald Gillespie's "Garden and Labyrinth of Time"
		text: "Now, the method consists herein: We made a box divided into various compartments: in which the" +
			" theoretical bases of all the sciences are set forth in tables in such a way that, " +
			"no matter what subject anyone may be asked to speak about, he will, " +
			"by various shiftings of tables and applications of a complex series of combinations soon find countless" +
			" arguments with which to elaborate any proposed question in all possible fullness through universal and" +
			" demonstrative syllogisms. Likewise, we have shown in our Musurgia (" +
			"method of composing music) how anyone, even if he has no knowledge of music, " +
			"can in the space of a single hour become capable of skillfully producing any melody you like.",
	},
	// {
	// 	author: "John Wilkins",
	// 	year: 1668,
	// 	title: "An Essay Towards a Real Character, and a Philosophical Language",
	// 	text: "And whereas several of the Species of Vegetables and Animals, " +
	// 		"do according to this present constitution, amount to more than Nine, " +
	// 		"in such cases the number of them is to be distributed into two or three Nines, " +
	// 		"which may be distinguished from one another by doubling the Stroke in some one or more parts of the" +
	// 		" Character; as supposed after this manner. If the first and most: simple Character we, be made use of, " +
	// 		"the Species that are fixed to it, will belong to the first combination of Nine; if the other, " +
	// 		"they will belong according to the order of them, unto the second Combination.",
	// },
	{
		author: "Aristotle",
		year: -335,
		title: "Poetics",
		text: "The subject I wish us to discuss is poetry itself, its species with their respective capabilities," +
			" the correct way of constructing plots so that the work turns out well, " +
			"the number and nature of the constituent elements of each species, " +
			"and anything else in the same field of inquiry.",
	},
}
