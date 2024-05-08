```
 █████╗ ██████╗ ██████╗ ██╗████████╗██████╗  █████╗  ██████╗ ███████╗    ██████╗  ██████╗  ██████╗  ██████╗ 
██╔══██╗██╔══██╗██╔══██╗██║╚══██╔══╝██╔══██╗██╔══██╗██╔════╝ ██╔════╝    ╚════██╗██╔═████╗██╔═████╗██╔═████╗
███████║██████╔╝██████╔╝██║   ██║   ██████╔╝███████║██║  ███╗█████╗       █████╔╝██║██╔██║██║██╔██║██║██╔██║
██╔══██║██╔══██╗██╔══██╗██║   ██║   ██╔══██╗██╔══██║██║   ██║██╔══╝       ╚═══██╗████╔╝██║████╔╝██║████╔╝██║
██║  ██║██║  ██║██████╔╝██║   ██║   ██║  ██║██║  ██║╚██████╔╝███████╗    ██████╔╝╚██████╔╝╚██████╔╝╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝    ╚═════╝  ╚═════╝  ╚═════╝  ╚═════╝
```

# Arbitrage3000

Arbitrage3000 is a simple Go application which finds arbitrage's between currencies, crypto, gold and silver.

## Requirements

To use A3k, you will need an API token from [OpenExchangeRates]([https://openexchangerates.org) and follow these steps:

1. Get your token from [OpenExchangeRates](https://openexchangerates.org)
2. Create a `.env` file in the root directory of the project
3. Add your token to the `.env` file:
```
OPENEXCHANGERATES_TOKEN=YOUR_TOKEN_HERE
```

## Usage
1. Clone this repo:
   ```
   git clone https://github.com/o6uoq/arbitrage3000.git
   ```

2. Build the Docker image:
   ```
   docker build -t arbitrage3000 .
   ```

3. Run the Docker image:
   ```
   docker run --env-file .env arbitrage3000
   ```

## ToDo
- [ ] Add a highlighted (red?) table for potentially discovered arbitrage's
- [ ] Add more information vis-a-vis gold and silver in the US vs UK, 3% discount for some purchases, 1% BitPay fees for others
- [ ] Either scrape websites for all live prices, or use an API (which is paid, scrapping is free)
- [ ] Backend API to create a friendly UI for this
    
