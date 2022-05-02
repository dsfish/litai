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
			"as I take it (it can also be called a general characteristic\nor algebra), " +
			"is that science in which are treated the forms or formulas of things in general,\nthat is, " +
			"quality in general or similarity and dissimilarity; in the same way that ever new\nformulas arise from" +
			" the elements a, b, c themselves when combined with each other, " +
			"whether\nthese elements represent quantities or something else.",
	},
	{
		author: "Thomas Dunn English",
		year: 1890,
		title: "Skeleton Essays, Or Authorship in Outline: Consisting of Condensed Treatises on Popular Subjects and" +
			" Directions How to Enlarge Them Into Essays, Or Expand Them Into Lectures",
		text: " Hence the preparation of this book, which will supply a need long existing, " +
			"and a demand\nwhich though moderate, is steady, and sure to increase. It presents a number of skeletons," +
			" or\noutlines, of essays on a variety of topics, " +
			"and in some instances these are followed by\ncompleted work, in order to show the reader how it is done.",
	},
	{
		author: "Gustav Freytag",
		year: 1900,
		title: "Technique of the Drama: An Exposition of Dramatic Composition and Art",
		text: "The poet of the present is inclined to look with amazement upon a method of work in which\nthe" +
			" structure of scenes, " +
			"the treatment of characters and the sequence of effects were\ngoverned by a transmitted code of fixed" +
			" technical rules. Such a limitation easily seems to us\nthe death of free artistic creation.",
	},
	{
		author: "Gottfried Leibniz",
		year: 1679,
		title: "Philosophical Papers and Letters",
		text: "Now, however, our characteristic will reduce the whole to numbers, " +
			"so that reasons can also\nbe weighted as if by a kind of statics for probabilities too, " +
			"will be treated in this calculation\nand demonstration, " +
			"since one can always estimate which of the given circumstances will\nmore probably occur. Finally, " +
			"anyone who is certainly convinced of the truth of religion and its\nconsequences, " +
			"and so embraces others in love that he desires this conversion of mankind.\nwill surely admit, " +
			"if he understands this matters, " +
			"that nothing will be more influential than this\ndiscovery for the propagation of the faith.",
	},
	{
		author: "Gustav Freytag",
		year: 1900,
		title: "Technique of the Drama: An Exposition of Dramatic Composition and Art",
		text: "If, however, we must deny ourselves the advantage of composing according to the\ncraftsman's" +
			" traditions which were peculiar to the dramatic art as well as the plastic art of\nformer centuries, " +
			"yet we should not scorn to seek, and intelligently to use, " +
			"the technical rules\nof ancient and modern times, which facilitate artistic effects of our stage.",
	},
	{
		author: "Aristotle",
		year: -335,
		title: "Poetics",
		text: "We therefore have (i) the mimesis of the action, the plot, " +
			"by which I mean the ordering of the\nparticular actions; (" +
			"i) the mimesis of the moral characters of the personages, " +
			"namely that in\nthe play which makes us say that the agents have certain moral qualities; (" +
			"ill) the mimesis of\ntheir intellect, " +
			"namely those parts in which they demonstrate something in speech or deliver\nthemselves of some general" +
			" maxim.",
	},
	{
		author: "Quirinus Kuhlmann",
		year: 1671,
		title: "Love-Kiss XLI",
		text: "From Night / Fog / Clash / Frost / Wind / Sea / Heat / South / East / West / North/ Sun." +
			"\nFire and Plagues /\nCome Day / Blaze / Bloom / Snow/ Peace / Land / Flash / Warmth / Heat / Joy/ Cool" +
			" /\nLight/ Flames and Dread/",
	},
	{
		author: "Georges Polti",
		year: 1916,
		title: "The Thirty-Six Dramatic Situations",
		text: "Which are the dramatic situations neglected by our own epic, " +
			"so faithful in repeating the few\nmost familiar? Which, on the other hand, " +
			"are most in use today? Which are the most\nneglected, and which the most used, in each epoch, genre, " +
			"school, author? What are the\nreasons for these preferences? The same questions may be asked before the" +
			" classes and\nsubclasses of the situations.",
	},
	{
		author: "Georges Polti",
		year: 1916,
		title: "The Thirty-Six Dramatic Situations",
		text: "It is equally natural that only the greatest and most complete civilizations should have\nevolved" +
			" their own particular conception of the drama, " +
			"and that one of these new conceptions\nshould be revealed by each new evolution of society, " +
			"whence arises the dim but faithful\nexpectation of our own age, " +
			"waiting for the manifestation of its own dramatic ideals, " +
			"before\nthe cenotaphs of an art which has long been, apparently for commercial reasons, " +
			"almost\nnon-existent.",
	},
	{
		author: "John Wilkins",
		year: 1668,
		title: "An Essay Towards a Real Character, and a Philosophical Language",
		text: "The second Part shall contain that which is the great foundation of the thing here designed," +
			"\nnamely a regular enumeration and description of all those things and notions, " +
			"to which marks\nor names ought to be assigned according to their respective natures, " +
			"which may be styled\nthe Scientific Part, comprehending Universal Philosophy. " +
			"It being the proper end and design\nof the several branches of Philosophy to reduce all things and" +
			" notions unto such a frame, as\nmay express their natural order, dependence, " +
			"and relations [..] The design of this Treatise\nbeing an attempt towards a new kind of Character and" +
			" Language, it cannot therefore be\nimproper to premise somewhat concerning those already in being; the" +
			" first Original of them,\ntheir several kinds, " +
			"the various changes and corruptions to which they are liable, " +
			"together\nwith the manifold defects belonging to them.",
	},
	{
		author: "Ramon Llull",
		year: 1308,
		title: "Ars Brevis",
		text: "And thus the intellect has a ladder for ascending and descending; as, for instance, " +
			"descending\nfrom a completely general principle to one neither completely general nor completely" +
			"\nparticular, and from a principle neither completely general nor completely particular to one\nthat is" +
			" completely particular. And in a similar fashion one can discuss the ascent of this\nladder ." +
			"] Everything that exists is implicit in the principles of this figure, " +
			"for everything is\neither good or great, etc., as God and angels, which are good, great, etc. Therefore," +
			" whatever\nexists is reducible to the above-mentioned principles.",
	},
	{
		author: "Ramon Llull",
		year: 1308,
		title: "Ars Brevis",
		text: "We have employed an alphabet in this Art so that it can be used to make figures, " +
			"as well as to\nmix principles and rules for the purpose of investigating the truth. For, " +
			"as a result of any one\nletter having many meanings, " +
			"the intellect becomes more general in its reception of the things signified, " +
			"as well as in acquiring knowledge. And this alphabet must be learned by\nheart, " +
			"for otherwise the artist will not be able to make proper use of this Art.",
	},
	{
		author: "Quirinus Kuhlmann",
		year: 1671,
		title: "Love-Kiss XLI",
		text: "At first glance, what is printed here above/seems impossible; yet just as certainly as two\ntimes two" +
			" makes six / so this too is the case. It might seem even more implausible/and here-\nI have the greatest" +
			" authorities to back me up / and can only argue this briefly here/though at\ngreater length elsewhere" +
			"/ that the major portion of Human Knowledge in fact lies hidden in\npermutation.",
	},
	{
		author: "Samuel Butler",
		year: 1872,
		title: "Erewhon",
		text: "Even a potato in a dark cellar has a certain low cunning about him which serves him in\nexcellent" +
			" stead. He knows perfectly well what he wants and how to get it. " +
			"He sees the light\ncoming from the cellar window and sends his shoots crawling straight thereto: they" +
			" will crawl\nalong the floor and up the wall and out at the cellar window, " +
			"if there be a little earth anywhere\non the journey he will find it and use it for his own ends.",
	},
	{
		author: "Athanasius Kircher",
		year: 1674,
		title: "Letter (from Gerald Gillespie's \"Garden and Labyrinth of Time\")",
		text: "Now, the method consists herein: We made a box divided into various compartments: in\nwhich the" +
			" theoretical bases of all the sciences are set forth in tables in such a way that, " +
			"no\nmatter what subject anyone may be asked to speak about, he will, " +
			"by various shiftings of\ntables and applications of a complex series of combinations soon find countless" +
			" arguments\nwith which to elaborate any proposed question in all possible fullness through universal and" +
			"\ndemonstrative syllogisms. Likewise, we have shown in our Musurgia (" +
			"method of composing\nmusic) how anyone, even if he has no knowledge of music, " +
			"can in the space of a single hour\nbecome capable of skillfully producing any melody you like.",
	},
	{
		author: "John Wilkins",
		year: 1668,
		title: "An Essay Towards a Real Character, and a Philosophical Language",
		text: "And whereas several of the Species of Vegetables and Animals, " +
			"do according to this present\nconstitution, amount to more than Nine, " +
			"in such cases the number of them is to be\ndistributed into two or three Nines, " +
			"which may be distinguished from one another by doubling\nthe Stroke in some one or more parts of the" +
			" Character; as supposed after this manner. If the\nfirst and most: simple Character we, be made use of, " +
			"the Species that are fixed to it, will\nbelong to the first combination of Nine; if the other, " +
			"they will belong according to the order of\nthem, unto the second Combination.",
	},
	{
		author: "Aristotle",
		year: -335,
		title: "Poetics",
		text: "The subject I wish us to discuss is poetry itself, its species with their respective capabilities," +
			"\nthe correct way of constructing plots so that the work turns out well, " +
			"the number and nature\nof the constituent elements of each species, " +
			"and anything else in the same field of inquiry.",
	},
}