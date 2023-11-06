package processline

import (
	"log"
	"os"
)

func CSIFileCreation() {
	data := []string{

		// HEADER RECORDS
		// CFH File Header
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + BSP Name(12 to 31) + Reserved Space (32 to 35) + Invoice Name (36 to 55) + Invoice Code (56 to 60) + Reserved Space(61)+ Processing Date (62 to 67) + Processing Time(68 to 71) + Billing Analysis Ending Date (72 to 77) + Handbook Revision Number(78 to 80) + Test/Production Status(81 to 84)+ ISO Country Code(85 to 86)+Reserved Space(87 to 256)

		// CIH Invoice Header
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Invoice Number (12 to 25) + Invoice Date (26 to 31) + Invoice Sequence Number(32 to 34) + Currency Type(35 to 38) + Reserved Space(38 to 44) + Airline/Credit Card Company Agreement Number(45 to 60) + Invoice Name(61 to 80) + Reserved Space(81 to 256)

		// CBH Batch Header
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Invoice Number (12 to 25) + Invoice Date (26 to 31) + Reserved Space (32 to 34) + Invoice Name (35 to 54) + Agent Numeric Code(55 to 62) + Batch Number (63 to 69) + Billing Analysis Ending Date(70 to 75) + Point of Sale Name (76 to 100) + Currency type (101 to 104) + Place of Issue (105 to 119) + Reserved Space(120 to 256)

		// TRANSACTION RECORDS
		// CBR Transaction Basic Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Invoice Number (12 to 25) + Invoice Date (26 to 31) + Billing Analysis Ending Date(32 to 37) + Agent Numeric Code(38 to 45) + Batch Number(46 to 52) + Credit Card Code(53 to 54) + Credit Card Account Number (55 to 73) + Expiry Date (74 to 77) + Approval Code (78 to 83) + Extended Payment Code(84 to 85)+ Ticket/Document Number(86 to 99) + Reserved Space(100) + Date of Issue(101 to 106) + Passenger Name(107 to 155) + Debit/Credit Code(156 to 157) + Form of Payment Amount (158 ti 168) + Authorised Amount (169 ti 179) + Customer File Reference(180 to 206) + Currency Type (207 to 210) +Flight Departure Date (211 to 215) + Transaction Code (216 to 219) + Reserved Space (220 to 230) + Related Ticket/Document Number (231 to 244) + Reserved Space (245) + Card Verification Value Result (246) + Statistical Code(247 to 249) + Reason for Issuance Code(250) + Routing Domestic/International Indicator (251) + Time of Issue (252 to 255) + Reserved Space (256)

		// COR Transaction Optional Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Billing Analysis Ending Date (12 to 17) + Ticket/Document Number (18 to 31) + Reserved Space (32) + Conjunction Ticket Indicator (33 to 35) + Origin Airport/City Code (36 to 40) + Destination Airport/City Code (41 to 45) + Stopover Code (46) + Reservation Booking Designator (47 to 48) + Carrier (49 to 51) + Reserved Space (52) + Fare Basis/Ticket Designator (53 to 67) + Origin Airport/City Code (68 to 72) + Destination Airport/City Code (73 to 77) + Stopover Code (78) + Reservation Booking Designator (79 to 80) + Carrier (81 to 83) + Reserved Space (84) + Fare Basis/Ticket Designator (85 to 99) + Origin Airport/City Code (100 to 104) + Destination Airport/City Code (105 to 109) + Stopover Code (110) + Reservation Booking Designator (111 to 112) + Carrier (113 to 115) + Reserved Space (116) + Fare Basis/Ticket Designator (117 to 131) + Origin Airport/City Code (132 to 136) + Destination Airport/City Code (137 to 141) + Stopover Code (142) + Reservation Booking Designator (143 to 144) + Carrier (145 to 147) + Reserved Space (148) + Fare Basis/Ticket Designator (149 to 163) + Reserved Space(164 to 166) + Source of Approval Code (167) + Form of Payment Transaction Identifier (168 to 192) + Passenger Specific Data (193 to 221) + Reserved Space (222 to 241) + Flight Departure Time (242 to 246) + Reserved Space (247 to 256)

		// COT Transaction Optional Tax Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Ticket/Document Number (12 to 25) + Reserved Space(26) + Tax/Miscellaneous Fee Type (27 to 34)  + Tax/Miscellaneous Fee Amount (35 to 48)+ Tax/Miscellaneous Fee Type (27 to 34)  + Tax/Miscellaneous Fee Amount (35 to 48)+ Tax/Miscellaneous Fee Type (49 to 56)  + Tax/Miscellaneous Fee Amount (57 to 70)	+ Tax/Miscellaneous Fee Type (71 to 78)  + Tax/Miscellaneous Fee Amount (79 to 92)+ Tax/Miscellaneous Fee Type (93 to 100)  + Tax/Miscellaneous Fee Amount (101 to 114)+ Tax/Miscellaneous Fee Type (115 to 122)  + Tax/Miscellaneous Fee Amount (123 to 136)+ Tax/Miscellaneous Fee Type (137 to 144)  + Tax/Miscellaneous Fee Amount (145 to 158)+ Tax/Miscellaneous Fee Type (159 to 166)  + Tax/Miscellaneous Fee Amount (167 to 180)	+ Tax/Miscellaneous Fee Type (181 to 188)  + Tax/Miscellaneous Fee Amount (189 to 202)+ Tax/Miscellaneous Fee Type (203 to 210)  + Tax/Miscellaneous Fee Amount (211 to 224)		+ Tax/Miscellaneous Fee Type (225 to 232)  + Tax/Miscellaneous Fee Amount (233 to 246) + Currency Type (247 to 250) + Reserved Space (251 to 256)

		// CAX 3DS and Additional Card Payment Information Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Ticket/Document Number (12 to 25) + Card Authentication Sequence Number (26 to 27) + 3D Secure Authentication and Additional Card Payment Information (28 to 126) + Reserved Space (127 to 256)

		// COE Electronic Miscellaneous Document Optional Coupon Detail Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Ticket/Document Number (12 to 25) + Conjunction Ticket Indicator (26 to 28) + EMD Coupon Number (29) + EMD Related Ticket/Document Number (30 to 43) + EMD Related Coupon Number (44) + EMD Service Type (45) + EMD Reason for Issuance Sub Code (46 to 48) + EMD Fee Owner Airline Designator (49 to 51) + EMD Operating Carrier (52 to 54) + EMD Attribute Group (55 to 57) + EMD Attribute Sub Group (58 to 60) + EMD Industry Carrier Indicator (61) + Reserved Space (62 to 256)

		//TOTAL RECORDS
		// CBT Batch Trailer Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Reserved Space (12 to 21) + Invoice Number (22 to 35) + Invoice Date (36 to 41) + Batch Number (42 to 48) + Billing Analysis Ending Date (49 to 54) + Total Debit Items (55 to 60) + Total Debit Amount (61 to 75) + Total Credit Items (76 to 81) + Total Credit Amount (82 to 96) + Currency Type (97 to 100) + Reserved Space (101 to 256)

		// CAT Totals Per Agent Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Agent Numeric code (12 to 19) + Billing Analysis Ending Date (20 to 25) + Invoice Number (26 to 39) + Invoice Date (40 to 45) + Total Debit Items Per Agent (46 to 51) + Total Debit Amount Per Agent (52 to 66) + Total Credit Items Per Agent (67 to 72) + Total Credit Amount Per Agent (73 to 87) + Agent First Batch Number (88 to 94) + Agent Last Batch Number (95 to 101) + Currency type (102 to 105) + Reserved Space (106 to 256)

		// CIT Invoice Trailer Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Invoice Number (12 to 25) + Invoice Date (25 to 31) + Invoice Sequence Number (32 to 34) + Airline/Credit Card Company Agreement Number (35 to 50) + Invoice Name (51 to 70) + Billing Analysis Ending Date (71 to 76) + Batch Count Per Invoice (77 to 81) + Total Debit Items Per Invoice (82 to 88) + Total Debit Amount Per Invoice (89 to 103) + Total Credit Items Per Invoice (104 to 110) + Total Credit Amount Per Invoice (111 to 125) + Currency type (126 to 129) + Total Discount Amount Per Invoice (130 to 144) + Total Tax on Discount Amount (145 to 159) + Total Net Amount Per Invoice (160 to 174) + Debit/Credit Code (175 to 176) + Reserved Space (177 to 256)

		//  CFT File Trailer Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Billing Analysis Ending Date (12 to 17) + Reserved Space (18 to 37) + Invoice Code (38 to 42) + Reserved Space (43) + Total Number of Batches (44 to 48) + Total Number of Invoices (49 to 52) + File Total Debit Amount (53 to 68) + File Total Credit Amount (69 to 82) + Reserved Space (83 to 256)

		// CEP Optional Extended Payment Plan Record
		// Standard Message Identifier(1 to 3) + Sequence Number(4 to 11) + Currency Type (12 to 15) + Extended Payment Total Amount (16 to 26) + Extended Payment Plan Code (27 to 32) + Extended Down Payment Amount (33 to 43) + Extended Payment Plan Instalment Amount (44 to 54) + Extended Payment Plan Instalment Quantity (55 to 56) + Payment Group Main Document (57 to 70) + Date of Issue Related Document (71 to 76) + Extended Payment Plan Original Invoice Number (77 to 90) + Ticket/Document Number (91 to 104) + Journey Turnaround Airport/City Code (105 to 109) + Extended Payment Taxes Amount (110 to 120) + Total Amount to Refund (121 to 131) + Reserved Space (132 to 256)

		"BFH0000000101GBE169230PROD2310190634BW004043",
		"BCH00000002021031231023D231018",
		"BOH000000030377200771231031BWP2",
	}

	// Create or open the file for writing
	file, err := os.Create("csifiles/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the data to the file
	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check for any errors during writing
	if err := file.Sync(); err != nil {
		log.Fatal(err)
	}

	// Success message
	log.Println("Data written to output.txt")
}
