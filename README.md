# Flipkart-Offers-And-Discount-API
# Flipkart Offers Service (Go + Gin + MongoDB)

This project provides APIs to store Flipkart offers and fetch the highest discount amount** based on bank and optionally by payment instrument.

---

## ✅ 1. Setup Instructions

### 1.1. Prerequisites
- [Go](https://golang.org/doc/install) (v1.19 or later recommended)  
- [MongoDB](https://www.mongodb.com/docs/manual/installation/) running locally or remotely  
- [Git](https://git-scm.com/)  

---

### 1.2. Install Dependencies
```bash
git clone <repository_url>
cd <project_folder>

go mod tidy
```

---

### 1.3. Configure Environment Variables
Create a `.env` file in the project root:

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB=flipkart_offers_db
SERVER_PORT=8080
```

---

### 1.4. Run the Server
```bash
go run main.go
```
Server will start at: `http://localhost:8080`

---

### 1.5. Postman Requests

#### 1. Save Offers (POST /offer)
Endpoint:  
```
POST http://localhost:8080/offer
```
Body (example):
```json
{
  "flipkartOfferApiResponse": {
    "offer_banners": [],
    "offer_sections": {
      "PBO": {
        "offers": [
          {
            "adjustment_id": "678",
            "summary": "Flat ₹1500 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹20000",
            "banks": ["HDFC"],
            "payment_modes": ["EMI_OPTIONS"],
            "emi_months": ["9", "12"],
            "image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
          }
        ]
      }
    }
  }
}
```
Response:
```json
{
  "noOfOffersIdentified": 1,
  "noOfNewOffersCreated": 1
}
```

---

#### 2. Get Highest Discount (GET /highest-discount)

✅ With payment instrument 
```
GET http://localhost:8080/highest-discount?amountToPay=20000&bank=HDFC&paymentInstrument=EMI_OPTIONS
```

✅ Without payment instrument (optional)
```
GET http://localhost:8080/highest-discount?amountToPay=20000&bank=HDFC
```

Response:
```json
{
  "highestDiscountAmount": 1500
}
```

---

## ✅ 2. Assumptions

1. Flipkart API response structure is consistent (offers are nested inside `offer_sections`).  
2. Discounts are stored in MongoDB in a flat structure (`banks`, `payment_modes`, `emi_months`).  
3. Only one bank and one payment instrument are used per query (scaling to multiple can be added later).  
4. Amount-based discount calculation is derived from `summary` parsing (e.g., "Flat ₹1500").  
5. If `paymentInstrument` is not provided, the highest discount is calculated based on the **bank only**.

---

## ✅ 3. Design Choices

1. Gin Framework  
   - Lightweight, fast, and idiomatic for Go REST APIs.  
   - Easy middleware & routing support.

2. MongoDB 
   - Flexible for unstructured Flipkart offer data.  
   - `banks` & `payment_modes` stored as arrays for quick `$in` queries.

3. Clean Architecture  
   - Controller → Service → Repository separation for testability.  
   - Models kept in a separate `models` package (no DB logic in services).

4. Logging  
   - Queries printed to console for debugging.  

---

## ✅ 4. Scaling the GET /highest-discount Endpoint (1,000 RPS)

1. Caching Layer
   - Use Redis to cache frequently accessed discount results per `(bank, paymentInstrument)` pair.

2. MongoDB Optimization
   - Add compound indexes on `{ banks, payment_modes }`.  
   - Use read replicas to distribute reads.

3. Horizontal Scaling 
   - Run multiple Gin instances behind a **load balancer** (e.g., Nginx or AWS ELB).

4. Async Pre-computation
   - Periodically compute highest discounts and store them in cache to avoid parsing & calculation on every request.

---

## ✅ 5. Future Improvements

1. Better Discount Parsing 
   - Implement NLP/regex-based parsing to support different discount formats (percentage, cashback).

2. Bulk Insert Optimization  
   - Use MongoDB `bulkWrite` to handle large Flipkart offer payloads efficiently.

3. Unit & Integration Tests
   - Add test coverage for service and repository layers.

4. Pagination & Filters 
   - Support pagination and filters (by bank, emi_months) for offer retrieval.

5. gRPC Support  
   - Expose a gRPC service for better performance if consumed internally.
   - i have worked with GRPC api's also Can handle them if required


Note:= There was no separate api for fetching the offers at flipkart instead i fetched the offers section from the item api from payments page whose screenshot was used in assignment pdf whose response was like this  

{
"response_status": "SUCCESS",
"messages": [],
"notify_messages": [
{
"title": "Please ensure your card can be used for online transactions. ",
"type": "TEXT"
},
{
"title": "Know More",
"type": "OVERLAY_LINK",
"url": "/api/v1/info/rbi"
}
],
"response_type": "PAYMENT_OPTIONS",
"token_version": "v3",
"disable_pay_button": false,
"disable_pay_timeout": 0,
"options": [
{
"applicable": true,
"selected": true,
"payment_instrument": "UPI",
"display_text": "UPI",
"messages": [
{
"type": "INFO",
"message": "Pay by any UPI app"
},
{
"type": "OFFER",
"message": "Up to ₹1000 instant saving on UPI"
}
],
"status_code": "",
"section": "OTHERS",
"priority": 0,
"provider": "FLIPKART",
"information": {
"logo_urls": {
"primary": "https://static-assets-web.flixcart.com/fk-p-linchpin-web/batman-returns/logos/UPI.gif"
}
}
},
{
"applicable": true,
"selected": false,
"payment_instrument": "CREDIT",
"display_text": "Credit / Debit / ATM Card",
"messages": [
{
"type": "INFO",
"message": "Add and secure cards as per RBI guidelines"
}
],
"status_code": "",
"section": "OTHERS",
"priority": 1,
"provider": "FLIPKART",
"information": {}
},
{
"applicable": true,
"selected": false,
"payment_instrument": "NET_OPTIONS",
"display_text": "Net Banking",
"messages": [
{
"type": "INFO",
"message": "This instrument has low success, use UPI or cards for better experience"
}
],
"status_code": "",
"section": "OTHERS",
"priority": 2,
"provider": "FLIPKART",
"information": {}
},
{
"applicable": true,
"selected": false,
"payment_instrument": "EMI_OPTIONS",
"display_text": "EMI (Easy Installments)",
"messages": [],
"status_code": "",
"section": "OTHERS",
"priority": 3,
"provider": "FLIPKART",
"information": {}
},
{
"applicable": true,
"selected": false,
"payment_instrument": "COD",
"display_text": "Cash on Delivery",
"messages": [
{
"type": "INTENT_INFO",
"message": "39,076 people used online payment options in the last hour. Pay online now for safe and contactless delivery.",
"icon": "https://rukminim1.flixcart.com/www/40/40/promos/09/12/2020/5309b678-8cf4-4290-8a66-761b0c8978b3.png?q=100"
}
],
"status_code": "",
"status_message": "",
"section": "OTHERS",
"priority": 4,
"information": {},
"callout": {}
},
{
"applicable": true,
"selected": false,
"payment_instrument": "EGV",
"display_text": "Gift Card",
"messages": [],
"status_code": "",
"status_message": "",
"section": "NEW_VOUCHERS",
"priority": 5,
"information": {},
"redeemed": 0,
"egv_payments": []
}
],
"price_summary": {
"remaining": 1906900,
"base_price": 1899000,
"item_count": 1,
"convertible_amount": [],
"price_details": [
{
"key": "PRICE",
"value": 1899000,
"item_count": 1,
"convertible_amount": []
}
],
"breakup": [
{
"key": "packaging_charges",
"display_text": "Packaging Charges",
"value": 0,
"type": "DEFAULT"
},
{
"key": "pickup_charges",
"display_text": "Pickup Charges",
"value": 0,
"type": "DEFAULT"
},
{
"key": "protect_promise_fee",
"display_text": "Protect Promise Fee",
"value": 7900,
"type": "DEFAULT"
}
],
"you_pay": [
{
"key": "AMOUNT_PAYABLE",
"value": 1906900,
"item_count": 0,
"convertible_amount": []
}
],
"notify_messages": []
},
"sla_summary": {
"prepaid": 1753034456430,
"postpaid": 1753034456430
},
"reservation_details": [
{
"reservationStatus": "RESERVED",
"presentTs": 1753034456430,
"ttl": 1753035296236,
"message": "You have time till July 20,2025,11:44:56 PM to complete your order."
}
],
"reservation_expiry_action": {
"url": "https://1.rome.api.flipkart.com/api/3/checkout/pgCancelResponse/desktop?redirect_domain=https://www.flipkart.com&callback=true",
"target": "https://1.rome.api.flipkart.com/api/3/checkout/pgCancelResponse/desktop?redirect_domain=https://www.flipkart.com&callback=true",
"action_type": "EXTERNAL_REDIRECTION",
"http_method": "POST",
"parameters": {
"reason_code": "RESERVATION_EXPIRED",
"merchant_transaction_id": "OD3349889764086141-TX-00",
"transaction_status": "FAILED"
}
},
"back_action": {
"url": "https://1.rome.api.flipkart.com/api/3/checkout/pgCancelResponse/desktop?redirect_domain=https://www.flipkart.com&callback=true",
"target": "https://1.rome.api.flipkart.com/api/3/checkout/pgCancelResponse/desktop?redirect_domain=https://www.flipkart.com&callback=true",
"action_type": "EXTERNAL_REDIRECTION",
"http_method": "POST",
"parameters": {
"reason_code": "CANCELLED_BY_USER",
"merchant_transaction_id": "OD3349889764086141-TX-00",
"transaction_status": "FAILED"
}
},
"offer_banners": [
{
"adjustment_sub_type": "NBFC_ZERO_INTEREST",
"adjustment_id": "FPO2502032024496ZX2L",
"summary": "No Cost EMI on Bajaj Finserv Cards",
"contributors": {
"payment_instrument": [],
"banks": [
"BAJAJFINSERV"
],
"emi_months": [
"3"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://static-assets-web.flixcart.com/apex-static/images/payments/banks/BFL_V2.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "EMI_FULL_INTEREST_WAIVER",
"adjustment_id": "FPO2502032024314IRYO",
"summary": "No Cost EMI on Credit Cards",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"AXIS",
"HDFC",
"ICICI",
"INDUSIND",
"KOTAK",
"SBI",
"PNB",
"FEDERALBANK",
"DBS",
"STANC",
"YESBANK",
"HSBC",
"RBL",x
"AMEX",
"FLIPKARTAXISBANK",
"BOBARODA"
],
"emi_months": [
"3",
"6",
"9",
"12"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250709162142RPSCI",
"summary": "Get additional ₹1000 off on Debit, Credit cards and UPI Transactions",
"contributors": {
"payment_instrument": [
"CREDIT",
"CREDIT",
"EMI_OPTIONS",
"NET_OPTIONS",
"PHONEPE",
"UPI_COLLECT",
"UPI_INTENT",
"FK_UPI"
],
"banks": [
"AXIS",
"HDFC",
"ICICI",
"INDUSIND",
"KOTAK",
"SBI",
"PNB",
"FEDERALBANK",
"DBS",
"STANC",
"YESBANK",
"HSBC",
"RBL",
"AMEX",
"FLIPKARTAXISBANK",
"BOBARODA"
],
"emi_months": [
"0",
"3",
"6",
"9",
"12"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250626172916JI4VZ",
"summary": "Flat ₹10 Cashback on Paytm UPI Trxns. Min Order Value ₹500. Valid once per Paytm account",
"contributors": {
"payment_instrument": [
"UPI_COLLECT",
"UPI_INTENT"
],
"banks": [],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250626182451RYQC7",
"summary": "5% off up to ₹750 on IDFC FIRST Power Women Platinum and Signature DebitCards. Min Txn Value: ₹5,000",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"IDFC"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/128/128/promos/01/09/2023/c00c8f1e-1ae7-49d1-b7e5-0e2df3d11ef7.png?q=100",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250619134128USHPF",
"summary": "5% cashback on Flipkart Axis Bank Credit Card upto ₹4,000 per statement quarter",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"FLIPKARTAXISBANK"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/axis-78501b36.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250619135528ZQIZW",
"summary": "5% cashback on Axis Bank Flipkart Debit Card up to ₹750",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"FLIPKARTAXISBANK"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/axis-78501b36.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250704192447CLSFR",
"summary": "Flat ₹750 on HDFC Bank Credit Card EMI on 6 months tenure. Min Txn Value: ₹15000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO2507042105097DZ77",
"summary": "Flat ₹500 on HDFC Bank Credit Card EMI on 6 months tenure. Min. Txn Value: ₹10000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO2506301938399SY1F",
"summary": "10% off up to ₹1,500 on BOBCARD EMI Transactions of 6 months and above tenures, Min TxnValue: ₹7,500",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"BOBARODA"
],
"emi_months": [
"6",
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/96/96/promos/31/07/2024/fb560aa6-1104-492b-b747-f458dbae837b.png?q=100",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO2507042000473JP81",
"summary": "Flat ₹1500 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹20000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO2507042306179OKXZ",
"summary": "Flat ₹1250 on HDFC Bank Credit Card EMI on 6 months tenure. Min Txn Value: ₹20000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250704194638XUV22",
"summary": "Flat ₹1000 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹15000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
},
{
"adjustment_sub_type": "PBO",
"adjustment_id": "FPO250704210925MIFUZ",
"summary": "Flat ₹750 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹10000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg",
"type": {
"value": "non-collapsable"
}
}
],
"offer_sections": {
"PBO": {
"title": "Partner offers",
"offers": [
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO250709162142RPSCI",
"summary": "Get additional ₹1000 off on Debit, Credit cards and UPI Transactions",
"contributors": {
"payment_instrument": [
"CREDIT",
"CREDIT",
"EMI_OPTIONS",
"NET_OPTIONS",
"PHONEPE",
"UPI_COLLECT",
"UPI_INTENT",
"FK_UPI"
],
"banks": [
"AXIS",
"HDFC",
"ICICI",
"INDUSIND",
"KOTAK",
"SBI",
"PNB",
"FEDERALBANK",
"DBS",
"STANC",
"YESBANK",
"HSBC",
"RBL",
"AMEX",
"FLIPKARTAXISBANK",
"BOBARODA"
],
"emi_months": [
"0",
"3",
"6",
"9",
"12"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png"
},
{
"adjustment_type": "CASHBACK_ON_CARD",
"adjustment_id": "FPO250626172916JI4VZ",
"summary": "Flat ₹10 Cashback on Paytm UPI Trxns. Min Order Value ₹500. Valid once per Paytm account",
"contributors": {
"payment_instrument": [
"UPI_COLLECT",
"UPI_INTENT"
],
"banks": [],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2502032024314IRYO",
"summary": "No Cost EMI on Credit Cards",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"AXIS",
"HDFC",
"ICICI",
"INDUSIND",
"KOTAK",
"SBI",
"PNB",
"FEDERALBANK",
"DBS",
"STANC",
"YESBANK",
"HSBC",
"RBL",
"AMEX",
"FLIPKARTAXISBANK",
"BOBARODA"
],
"emi_months": [
"3",
"6",
"9",
"12"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/64/64/promos/02/06/2022/a89d0cb0-9155-4545-bd47-c5c53c8d50b7.png"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO250626182451RYQC7",
"summary": "5% off up to ₹750 on IDFC FIRST Power Women Platinum and Signature DebitCards. Min Txn Value: ₹5,000",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"IDFC"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/128/128/promos/01/09/2023/c00c8f1e-1ae7-49d1-b7e5-0e2df3d11ef7.png?q=100"
},
{
"adjustment_type": "CASHBACK_ON_CARD",
"adjustment_id": "FPO250619134128USHPF",
"summary": "5% cashback on Flipkart Axis Bank Credit Card upto ₹4,000 per statement quarter",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"FLIPKARTAXISBANK"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/axis-78501b36.svg"
},
{
"adjustment_type": "CASHBACK_ON_CARD",
"adjustment_id": "FPO250619135528ZQIZW",
"summary": "5% cashback on Axis Bank Flipkart Debit Card up to ₹750",
"contributors": {
"payment_instrument": [
"CREDIT"
],
"banks": [
"FLIPKARTAXISBANK"
],
"emi_months": [
"0"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/axis-78501b36.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO250704192447CLSFR",
"summary": "Flat ₹750 on HDFC Bank Credit Card EMI on 6 months tenure. Min Txn Value: ₹15000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2507042105097DZ77",
"summary": "Flat ₹500 on HDFC Bank Credit Card EMI on 6 months tenure. Min. Txn Value: ₹10000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2506301938399SY1F",
"summary": "10% off up to ₹1,500 on BOBCARD EMI Transactions of 6 months and above tenures, Min TxnValue: ₹7,500",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"BOBARODA"
],
"emi_months": [
"6",
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://rukminim1.flixcart.com/www/96/96/promos/31/07/2024/fb560aa6-1104-492b-b747-f458dbae837b.png?q=100"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2507042000473JP81",
"summary": "Flat ₹1500 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹20000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2507042306179OKXZ",
"summary": "Flat ₹1250 on HDFC Bank Credit Card EMI on 6 months tenure. Min Txn Value: ₹20000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"6"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO250704194638XUV22",
"summary": "Flat ₹1000 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹15000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO250704210925MIFUZ",
"summary": "Flat ₹750 on HDFC Bank Credit Card EMI on 9 months and above tenure. Min Txn Value: ₹10000",
"contributors": {
"payment_instrument": [
"EMI_OPTIONS"
],
"banks": [
"HDFC"
],
"emi_months": [
"9",
"12",
"18",
"24",
"36"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://img1a.flixcart.com/www/linchpin/fk-cp-pay/hdfc-fa424ee4.svg"
},
{
"adjustment_type": "INSTANT_DISCOUNT",
"adjustment_id": "FPO2502032024496ZX2L",
"summary": "No Cost EMI on Bajaj Finserv Cards",
"contributors": {
"payment_instrument": [],
"banks": [
"BAJAJFINSERV"
],
"emi_months": [
"3"
],
"card_networks": []
},
"display_tags": [
"PAYMENT_OPTIONS"
],
"image": "https://static-assets-web.flixcart.com/apex-static/images/payments/banks/BFL_V2.svg"
}
]
}
},
"merchant_id": "mp_flipkart",
"section_details": {
"PREFERRED": {
"visible_options_count": 3
},
"OTHERS": {
"visible_options_count": 10
},
"NEW_VOUCHERS": {
"visible_options_count": 1
},
"LINKED_VOUCHERS": {
"visible_options_count": 1
}
}
}

