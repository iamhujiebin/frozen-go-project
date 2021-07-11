package main

import "fmt"

func main() {
	greetBlacks := []int64{240420, 302265, 108916, 184756, 162539, 216409, 154045, 138856, 303796, 157636,
		144681, 103350, 302449, 289647, 156827, 165073, 100487, 305102, 281689, 302384, 100436, 304035, 302512,
		104978, 195905, 290755, 116062, 274360, 108271, 100613, 154135, 128895, 171376, 168125, 156287, 141071, 240541,
		100430, 101421, 161391, 124298, 195110, 112956, 302480, 302380, 100425, 100493, 100386, 102252, 190584, 241003, 305390,
		142320, 149821, 113407, 100531, 100803, 190447, 100538, 129891, 102141, 106570, 215109, 302515, 101519, 263915, 100439,
		162352, 102258, 118170, 303722, 302419, 179298, 156182, 124753, 249317, 302438, 238346, 169596, 275004, 113498, 249223, 105847,
		239453, 134102, 236903, 169857, 308173, 156257, 181403, 154007, 100384, 109795, 120847, 100480, 238436, 281687, 151940,
		302495, 267423, 248837, 156318, 305410, 264407, 108730, 100423, 292467, 136484, 272469, 100263, 100545, 303698, 194959,
		295474, 259515, 104177, 302415, 246684, 107715, 160910, 154002, 182063, 103381, 102692, 249264, 300795, 183190,
		117957, 257228, 112821, 100512, 300950, 228360, 169655, 113287, 157690, 137197, 249268, 196320, 266349, 285217, 216712,
		167482, 100465, 101438, 259477, 152855, 219110, 167074, 132887, 160661, 197865, 263402, 190861, 100435, 204064, 183005,
		101143, 149124, 100376, 100375, 267286, 159522, 214842, 231267, 182088, 105003, 103425, 102005, 102246, 181486,
		101437, 162509, 240325, 147315, 182359, 100207, 263537, 248101, 197735, 148424, 168178, 100371, 130728, 113062, 101442,
		320689, 308340, 326230, 319122, 103009, 101469, 308223, 327587, 304734, 305418, 117957, 305102, 316944, 205581, 306372,
		195110, 274360, 320803, 302448, 308204, 322762, 190679, 156827, 304631, 312633, 118782, 101164, 105358, 281689, 156426,
		329472, 240420, 102276, 318702, 313506, 154203, 276115, 100682, 118047, 100391, 310930, 112175, 325018, 100403, 306747,
		324369, 319049, 267084, 308250, 313599, 302384, 302384, 102281, 154480, 113270, 213500, 143052, 157690, 100515, 103370,
		249502, 324095, 313359, 325161, 329358, 100263, 325359, 322540, 308426, 326334, 243441, 100278, 103425, 128117, 102185,
		100618, 263537, 105239, 102192, 248837, 133231, 106618, 141638, 100474, 103314, 190447, 204150, 304745, 101464, 304886,
		300795, 313241, 118148, 154002, 102894, 181403, 231453, 307825, 153171, 138176, 100368, 148943, 319460, 155453, 122324,
		316958, 328549, 100458, 104624, 171482, 293103, 311970, 101437, 153964, 240534, 308949, 100243, 325461, 160910, 184756,
		275047, 313268, 318400, 120365, 154135, 162352, 184200, 281687, 308620, 310405, 316185, 320421, 116948, 156257, 178564,
		245195, 275837, 302535, 308144, 313493, 320649, 324360, 329397, 100429, 101477, 113407, 116839, 124753, 130728, 157636,
		224056, 238436, 267011, 305214, 305410, 308310, 313767, 318709, 319301, 319502, 323917, 325601, 325954, 326046, 100240,
		100747, 322058, 100001, 100002, 100205, 100207, 100209, 100212, 100214, 100216, 100217, 100219, 100231, 100232, 100235,
		100242, 100245, 100249, 100250, 100254, 100256, 100257, 100259, 100260, 100261, 100264, 100265, 100267, 100269, 100272,
		100274, 100275, 100276, 100279, 100280, 100281, 100282, 100283, 100284, 100285, 100286, 100287, 100289, 100290, 100292,
		100293, 100294, 100308, 100309, 100313, 100316, 100417, 100418, 100419, 100421, 100423, 100424, 100425, 100426, 100428,
		100430, 100431, 100434, 100435, 100436, 100437, 100438, 100439, 100450, 100528, 100529, 100530, 100531, 100534, 100535,
		348689, 306747, 350534, 101164, 154673, 348702, 340718, 343896, 117957, 305527, 103370, 308310, 322540, 348538, 313268,
		350411, 103425, 100278, 105239, 248837, 100618, 138756, 332737, 343898, 141638, 100517, 309763, 337036, 319049, 335482,
		154480, 181403, 300795, 318400, 154203, 190447, 341945, 328310, 195110, 100240, 100747, 231453, 325359, 101464, 102192,
		103009, 327136, 336946, 304745, 305648, 350752, 100375, 100458, 112175, 118148, 168263, 190679, 195905, 240541, 297707,
		304886, 316944, 325018, 326230, 327587, 100231, 100393, 100561, 108916, 113270, 118047, 118170, 142436, 147167, 156257,
		228929, 267084, 275047, 308040, 308426, 310748, 313359, 319121, 320649, 323917, 324369, 324385, 340859, 348085, 348117,
		350142, 350324, 102894, 157690, 263537, 100001, 243441, 313506, 308620, 351600, 348371, 352630, 353710, 104658, 141638, 351378, 348117,
		348051, 351362, 235132, 305527, 112175, 327589, 351996, 156372, 353288, 117957, 353714, 276115, 240534, 313599, 343893, 156257, 350534,
		352194, 348738, 342115, 340718, 322540, 350274, 148943, 350752, 340689, 350411, 362473, 350870, 313767, 164607, 213500, 351362, 362291, 276115, 353379,
		320689, 356344, 358955, 356511, 353834, 353714, 350274, 362039, 362253, 340718, 100383, 352194, 357039, 350752, 351996, 155453, 184756, 361944, 358544, 352818,
		353859, 350324, 353469, 350534, 356020, 361886, 361965, 355937, 235132, 310405, 364910, 181403, 264078, 316915, 351600, 364522, 103425, 100278, 105239, 100618,
		353467, 100376, 351461, 304886, 169045, 334216, 349887, 356821, 153668, 102894, 356082, 141638, 298106, 327136, 231453, 181486, 358496, 100747, 149124, 356228,
		281687, 319049, 335482, 195110, 355719, 204064, 156182, 304745, 361514, 195905, 190679, 306372, 320421, 320649, 353297, 249571, 304810, 308340, 319122, 325779,
		350283, 353227, 362240, 297707, 313359, 348927, 358562, 104958, 133231, 100375, 102281, 103009, 103502, 117842, 118047, 118170, 264407, 275837, 298433, 328549,
		342400, 343898, 348371, 348538, 352630, 353066, 357691}
	m := make(map[int64]struct{})
	for _, v := range greetBlacks {
		m[v] = struct{}{}
	}
	fmt.Printf("a:%v,b:%v\n", len(greetBlacks), len(m))
}
