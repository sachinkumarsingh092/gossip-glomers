==Phrack Inc.==
[Volume 0x10, Issue 0x47, Phile #0x11 of 0x11](http://phrack.org/issues/71/17.html#article)

|=-----------------------------------------------------------------------=|
|=-----------------------=[ Calling All Hackers ]=-----------------------=|
|=-----------------------------------------------------------------------=|
|=--------------------------=[ cts (@gf_256) ]=--------------------------=|
|=-----------------------------------------------------------------------=|

### Table of Contents

0 - Preamble
1 - About the Author
2 - The Birth of a Shitcoin
3 - How Money Works
    3.1 - Fixed Income
    3.2 - Equities
    3.3 - Shareholder Value
4 - Startup Blues
5 - Takeaways
6 - Thanks
7 - References
8 - Appendix

### 0 - Preamble

Hi. 

I'm cts, also known as gf_256, ephemeral, or a number of other handles.
I am a hacker and now a small business owner and CEO. In this article, 
I would like to share my experience walking these two different paths.

A hacker is someone who understands how the world works. It's about 
knowing what happens when you type "google.com" and press Enter. It's 
about knowing how your computer turns on, about memory training, A20, 
all of that. It's about modern processors, their caches, and their side 
channels. It's about DSi bootloaders and how the right electromagnetic 
faults can be used to jailbreak them. And it's about how Spotify and 
Widevine and AES and SGX work so you can free your music from the 
shackles of DRM.

But being a hacker is so much more than these things. It's about knowing 
where to find things. Like libgen and Sci-Hub and nyaa. Or where to get 
into the latest IDA Pro group buy. Or which trackers have what and how 
to get into them.

It's about knowing how to bypass email verification. How to bypass SMS 
verification. How to bypass that stupid fucking verification where you 
hold your driver's license up to a webcam (thank you, OBS virtual camera!) 
Having an actual threat model not just paranoia. Knowing that you're not 
worth burning a 0day on, but reading indictments to learn from others' 
mistakes.

It's about knowing where to buy estradiol valerate on the internet and how 
to compound injections. Or the "bodybuilder method" to order your own 
blood tests when your state requires a script to do so. It's about knowing 
which shipments give the US CBP a bad vibe and which don't.

It's about knowing what happens when you open Robinhood and giga long NVDA 
FDs. I mean the actual market microstructure, not "Ken Griffin PFOF bad". 
Then using that microstructure to find an infinite money glitch (high 
Sharpe!). It's about knowing how to get extra passports and reading the
tax code. 

It's about knowing how to negotiate your salary (or equity). It's about 
knowing why things at the supermarket cost what they do. Or how that awful 
shitcoin keeps pumping. And why that dogshit startup got assigned that 
insane valuation. And understanding who really pays for it in the end 
(hint: it's you).

My point is, it is not just about computers. It's about understanding how 
the world works. The world is made up of people. As much as machines keep 
society running, those machines are programmed by people--people with 
managers, spouses, and children; with wants, needs, and dreams. And it is 
about using that knowledge to bring about the change you want to see.

That is what being a hacker is all about.


### 1 - About the Author

I have been a hacker for 13 years. Prior to founding Zellic, I helped
start a CTF team called perfect blue (lately Blue Water). We later became
the number one ranked CTF team in the world. We've played in DEF CON CTF.
We've won GoogleCTF, PlaidCTF, and HITCON. It's like that scene from
Mr. Robot but not cringe.

In 2021, we decided to take that hacker friend circle and form a security 
firm. It turned out that crypto paid well, so we worked with a lot of 
crypto clients. In the process, we encountered insane, hilarious, and 
depressingly sobering bullshit. In this article, I will tell some stories 
about what that bullshit taught me, so you can benefit from the same 
lessons as I have.

Markets are computers; they compute prices, valuations, and the allocation 
of resources in our society. Hackers are good at computers. Let's learn 
more about it.


### 2 - The Birth of a Shitcoin

I can't think of a better example than shitcoins. Let's look at the
crypto markets in action.

First, let's talk about tokens. What is their purpose? The purpose of a 
token is to go up. There is no other purpose. Token go up. This is 
important, remember this point.

Now the question is, how do we make the token go up? In crypto, there are 
two main kinds of token deals. Let's call them the Asian Arrangement and 
the Western Way.

The Asian Arrangement is a fairly straightforward pump and dump. It's a 
rectangle between the VC, the Market Maker, the Crypto Exchange, and the 
Token Project Founder.

1. The exchange's job is to list the token, bringing in investors. They 
   get paid in a mix of tokens and cold, hard cash. Their superpower is 
   owning the customer relationships with the retail users, and the 
   naming rights to sports arenas.

2. The market maker provides liquidity so the market looks really 
   healthy and well-traded so it is easy to buy the token. In good 
   deals, they are paid in in-the-money call options on the tokens, 
   so they are incentivized to help the token trade well. Their 
   superpower is having a lot of liquidity to deploy, and people 
   on PagerDuty.

3. The founder's job is to pump the token and shill it on Twitter. 
   They are the hype man, and it's their job to drum up the narrative 
   and pump everyone's bags. Their unique power is they can print more 
   tokens out of thin air, and this is in large part how they get paid 
   in this arrangement.

4. Lastly, the VC gets paid to organize the deal. They give the founders 
   some money, who in return give a pinky promise that they will give 
   the VC a lot of tokens once the tokens actually exist. This is known 
   as a Simple Agreement for Future Tokens, or SAFT. Their superpower is 
   dressing up the founders and project so it seems like the Next Big 
   Thing instead of a Ponzi scheme.

Everyone gets paid a ton of token exposure (directly or indirectly), 
and when it lists, it pumps. Then the insiders dump and leave with a 
fat stack. Except retail, they end up with the bag. 

Sometimes the listing doesn't go well for the organizers, in which case, 
better luck next time. But retail always loses.

```
  wtf???   LFG!!! to the moon   
       ,o  \oXo/\o/          
       /v   | |  |
      /\   / X\ / \

    crypto investors
        ^ |
        | |
        | v
    +----------+                provides liquidity          +--------+
    |  Crypto  |  <---------------------------------------  | Market |
    | Exchange |  ----------------------------------------> | Maker  |
    +----------+                   maker fees               +--------+
        ^ |                                                    ^     
  fees, | | listing                                    options |     
 tokens | |                                            / fees  |      
        | |  +-------------------------------------------------+
        | v  |                                                        
    +---------+       tokens / SAFT / token warrants       +---------+
    |  Token  |  --------------------------------------->  | Venture |
    | Project |  <---------------------------------------  | Capital |
    +---------+     cash , intros to CEX / MM, shilling    +---------+
```

This machine worked exceptionally well in 2017, especially before China 
banned crypto. All those ICO shitcoins? Asian Arrangement. And it still 
works well to this day, except people are more wary of lockups and vesting 
schedules and so on.

Now let's discuss the Western Way. The Asian Arrangement? That old pump 
and dump? No sir, we are civilized people. Instead, our VCs *add value* 
to their investments by telling the world "how disruptive the tech is" 
and how the "team are incredible outliers". And they will not blatantly 
PnD the token, but instead they will fund "projects in the ecosystem" so 
it appears there is real activity happening on the platform. 

This is to hype up metrics (like TPS or TVL) to inflate the next round 
valuation. Anyways, then they dump. Or maybe the VC is also a market 
maker so they market make their portfolio company tokens. Overall it's 
the same shit (Ponzi) but dressed up in a nicer outfit.

Asian Arrangement or Western Way--either way, if you're the token founder, 
your main priority is to just GO TO MARKET NOW and LAUNCH THE TOKEN. This 
is so you can collect your sweet bag and dump some secondary before 
someone else steals the narrative or the hype cycle moves on.

This is one of the reasons there are so many hacks in crypto. The code is 
all shitty because it's rushed out as fast as possible by 20-something-
year-old software engineers formerly writing Typescript and Golang at
Google. Pair that with some psycho CEO product manager. Remember, it is
not about WRITING SECURE CODE, it is about SHIPPING THE FUCKING PRODUCT.
Good luck rewriting it in Rust!

All of this worked well until Luna, then 3AC, Genesis, and FTX imploded in 
2022. It still works, but you have to be less blatant now.

Shitcoins do serve an essential need. They are an answer to financial 
nihilism. Many people are working dead-end wage slave jobs that are not 
enough to "make it". They feel trapped and forced to work at jobs they 
fucking hate and waste their life doing pointless shit to generate 
shareholder value. This kind of life feels unacceptable, yet there are 
few avenues out. So what is the only "attainable" solution left? Gamble 
it on shitcoins, and if you lose...maybe next paycheck will be better.

But enough about crypto, let's talk about securities.


### 3 - How Money Works

#### 3.1 - Fixed Income

First, let's start with fixed income. I'm talking boring, old-fashioned
bonds, like Treasury bonds. A lot of people are introduced nowadays to 
finance through equities (stocks) and tokens. In my opinion, this is 
only half of the story. Fixed income is the bedrock of finance. It has 
fundamental value. It provides a prototypical asset that all assets can 
be benchmarked based on.

Fixed income assets, like bonds, boil down to borrowing and lending. A 
bond is basically an IOU for someone to pay you in the future. It is more 
useful to have a dollar today than in a year, so lenders charge a fee for 
access to money today. This fee is known as interest, and how it is baked 
into the equation varies from asset-to-asset. Some bonds come with 
interest payments, whereas other bonds are zero-coupon. The most important 
thing is to remember that bonds are essentially an IOU to pay $X in the 
future.

Here is an example. Let's say you would like to borrow $100 to finance an 
upcoming project. The interest rate will be 5% per year. To borrow money, 
you would issue (mint) a bond (an IOU) for $X+5 dollars to be repaid 1 
year in the future. In exchange for this fresh IOU, the lender will give 
you $X dollars now. 

On the lender's balance sheet, they will be less $X dollars worth of cash, 
but will also have gained ($X+5) dollars worth of an asset (your IOU), 
creating $5 of equity. In contrast, you would have $X more cash in assets, 
but also an ($X+5) liability, creating -$5 of equity. 

This example also works for depositing money at a bank. Here, you are the 
lender, and the bank is the borrower. Your deposits would be liabilities 
on their balance sheet, as they are liable to pay you back the deposit if 
you choose to withdraw it.

```
     Lender's Balance Sheet               Borrower's Balance Sheet   
   ===========================          ===========================  
    Assets:                              Assets:
      IOU-----------------X+5              Cash------------------X
                                         
    Liabilities:                         Liabilities:
      Cash----------------(X)              IOU-----------------X+5
                                         
    Equity:                              Equity:
      Equity----------------5              Equity--------------(5)
```

Fixed income assets are extremely simple. There are various risks (credit 
risk, interest rate risk, etc.), but excluding these factors, you 
essentially get what you pay for. Unlike a token or stock, the bond is not 
going to suddenly evaporate or crash. (In theory.) Because of this, they
can be modeled in a straightforward way; a way so straightforward even
a high school student can understand it.

Let's say I have $X today. Suppose the prevailing (risk-free) interest 
rate is 5%. What is the value of this $X in a year? Obviously, it would be 
no less than $X*1.05, as I can just lend it out for 5% interest and get 
$X*1.05 back in a year. If you gave me the opportunity to invest in any 
asset yielding less than 5%, this would be a bad deal for me, since I 
could just lend it out myself to get 5% yield.

Now, let's analyze the same scenario, but in reverse. Let's take that IOU 
from earlier. What is the value *today* of a (risk-free) $X IOU, due in 1 
year? It would be worth no more than $X/1.05. This is because with $X/1.05 
dollars today, I could lend it out and collect 5% interest to end up with 
$X again in the future. If I pay more than $X/1.05, I am getting a bad 
deal, since I am locking up my money with you when it would be more 
capital efficient to just lend it out myself.

You can probably see where I am going with this. The present value of an 
$X IOU at some time *t* in the future is $X/(1+r)^t, where *r* is the 
discount rate. The discount rate describes the "decay" of the value over 
time, due to interest but also factors like potential failure of the asset 
(for example, if the asset is a company, business failure of the company). 

Now, if we have some asset which pays a series of future cash flows 
*f(t)*, we can model this asset as a bundle of IOUs with values f(t) due 
in time 1, 2, 3, and so on. Then the present value of this asset is the 
geometric series sum of the discounted future cash flows. This is called 
discounted cash flows (DCF). Congrats, now you can do better modeling than 
what goes into many early-stage venture deals.

```
   +------+-----+-----+---------+---------+---------+-------+---------+
   | Year |  0  |  1  |    2    |    3    |    4    |  ...  |    t    |
   +------+-----+-----+---------+---------+---------+-------+---------+
   | Cash | CF1 | CF2 |   CF3   |   CF4   |   CF5   |  ...  |  CF_t   |
   | Flow |     |     |         |         |         |       |         |
   +------+-----+-----+---------+---------+---------+-------+---------+
   | Disc.| CF1 |_CF2_| __CF3__ | __CF4__ | __CF5__ |  ...  | _CF_t__ |
   | Val  |     | 1+r | (1+r)^2 | (1+r)^3 | (1+r)^4 |       | (1+r)^t |
   +------+-----------+---------+---------+---------+-------+---------+
           IOU 1 IOU 2   IOU 3     IOU 4     IOU 5     ...     IOU n

         inf
          _   f(t)                                               1
   DCF = \  ------- = (assume constant annual cash flow x) = --------- x
         /_ (1+r)^t                                          1-1/(1+r)
         t=0
   
       = (1/r + 1) x
   
   Cash flow multiple = (value) / (annual cash flow) ~= 1/r
```

(The astute reader might also find that they can go backwards from 
valuations to estimate first, second, ... Nth derivatives of the cash 
flow or the year-to-year survival chances of a company. And these can be 
compared with...going outside and touching grass to see if the valuation 
actually makes sense.)

At this point, you're probably wondering why I'm boring you with all of 
this dry quant finance 101 shit. Well, it's a useful thing to know about 
how the world works.

First, interest rates affect you directly and personally. You may have 
heard of the term "zero interest rate environment". In a low interest rate 
environment, cash flow becomes irrelevant. Why? Consider the DCF geometric 
series sum if the interest rate r = 0. The present value approaches 
infinity. If the benchmark hurdle rate we're trying to beat is 0%, 
literally ANYTHING is a better investment than holding onto cash. 

Now do you see why VCs were slamming hundreds of millions into blatantly 
bad deals and shit companies during Covid? Cash flow and profitability 
didn't matter, because you could simply borrow more money from the money 
printer.

Here's a more concrete example. Do you remember a few years ago when Uber 
rides were so cheap, that they were clearly losing money on each ride? 

This is known as Customer Acquisition Cost, or CAC. CAC is basically the 
company paying you to use their app, go to their store, subscribe to the 
thing, ... whatever. The strategy is well-known: burn money to acquire 
users until everyone else dies and you become a monopoly. Then raise the 
prices. 

But here is the key point: this only works in a low-interest rate 
environment. In such an environment, discounting is low, and thus, future 
growth potential is valued over profitability and fundamentals at present. 
It doesn't need to make sense *today* as long as it works 10 years from 
now. For now, we can keep borrowing more money to sustain the burn.

Of course, when rates go back up, the free money machine turns off and 
the effects ripple outward. You are the humble CAC farmer, farming CAC 
from various unprofitable consumer apps like ride share, food delivery, 
whatever. These apps raise their money from their investors, VC and 
growth equity funds. These funds in turn raise their money from *their* 
investors, their limited partners. These LPs might be institutional 
capital like pension funds, sovereign wealth funds, or family offices. 

At the end of the day, all of that wealth is generated somewhere 
throughout the economy by ordinary people. So when some VC-backed 
founders throw an extravagant party on a boat with fundraised dollars, 
in some sense, you are the one paying for it.

And when the money machine turns off, anyone who had gotten complacent 
under ZIRP is now left scrambling. Companies will overhire during ZIRP 
only to do layoffs when rates go up.

```
                         +=========================+                       
                         |   THE LIQUIDITY CYCLE   |                       
                         +=========================+                       
                                                                           
                                                                           
                                             VENTURE CAPITAL               
                   _______________      ,.-^=^=^=^=^=^=^=^=^=^;,           
                 ,;===============>>   E^ a16z   LSVP    Tiger '^3.        
               .;^                    E^       FF    Social Cap. '^3       
              //  condensation       .E    Bain   SoftBank  Accel 3^       
             /|^                     ^E  KP          Benchmark    :^       
             ||                       ^;:   YC    Greylock   GC  ;3'       
     ,.^-^-^-^-^-^-^-^-^-^-^;,          ^.=.=_=_=_=_=_=_=_=_=_=_=^         
    E^ endowments    family '^:.            \\\\\\\\\\\\\\\\\\\\           
   E^                offices  '^3            \\\\\\\\\\\\\\\\\\\\          
  E'  pension                  ^3. SOURCE     \\\ precipitation \\         
  ^;   funds       sovereign   3.' CAPITAL     \\\\\\\\\\\\\\\\\\\\        
   E;:           wealth funds ,3^  (LPs)        \\\\\\\\\\\\\\\\\\\\       
    ^;._.._._._._._._._._._._,^                  \\\\\\\\\\\\\\\\\\\\      
                                                               /\          
      ^ ^ ^ ^ ^ ^ ^ ^                      gamefi   /\  /\  uber eats      
      | | | | | | | |                     shitcoins/::\/::\  /::::\   /\   
      | evaporation |                             / doordash/^^^^^^\ /^^\  
      | | | | | | | |         ____________       /      \  /     hello   \ 
                             (poggers desu)     /_____ lime ____ fresh ___\
    \o/ \oXo/\oXoXo/  o       '=========='       UNPROFITABLE CONSUMER APPS
     |   | |  | | |  /|\         Oo._ /\_/\                 ,///           
  __/_\_/_X_\/_X_X_\_/_\__ /_________(@'w'@)_____________.,://'            
          SOCIETY          \''''''''  -...-''''''''''''''''' surface       
                                    THE HUMBLE               runoff        
                                    CAC FARMER                             
```

Second, credit is not inherently a bad thing if used responsibly. Take for 
example those Buy Now, Pay Later loans. Now that you are equipped with the 
concept of capital efficiency, wouldn't it technically better than paying 
cash to take an interest-free BNPL loan and temporarily stick the freed 
cash into an investment? (Barring other side effects, etc.)

Third, the concept of net present value--i.e., credit--is the killer app 
of finance. It allows you to transport value from the future into today. 
Of course, that debt must be repaid in the future, unless you can figure 
out a way to kick the can down the road forever.

For now, let's get back to stocks.

#### 3.2 - Equities

Now we have seen both sides of the coin. Asset value is twofold: 
speculative and fundamental.

First, we saw speculative value as illustrated by crypto meme coins. Then, 
on the other hand, we examined fundamental value as illustrated by, e.g. a 
US Treasury. These two lie on two extremes of a spectrum. Some sectors and 
stocks are more speculative than others; Nvidia is practically a meme coin 
at this point, whereas something like Coca-Cola is like fixed income for 
boomers (NFA BTW). Most assets have a blend of both.

Thinking about stocks, they (usually) have some fundamental value. 
Equities represent ownership of some asset, like a business. The business 
in theory generates dividends for shareholders, and this cash flow (or the 
net present value of future ones) represents the fundamental value of the 
business. As we've seen, assets with better cash flows are more valuable.

In practice, buybacks can be used to create what is effectively a 
shareholder dividend in a more tax-advantaged way. Whereas with dividends, 
they are taxed as income, and this is realized immediately. With buybacks, 
they are taxed as capital gains, but crucially the gains are not realized 
until the asset is sold. This could be indefinitely far in the future, so 
it's more capital efficient. It has the added benefit that it helps pump 
the token, and imo this is kind of cute because it marries both the 
fundamental and speculative aspects.

Meanwhile, like tokens, stocks are also supposed to go up. Here's an 
example: imagine a generic meme coin. Apart from Go Up, what does it do? 
Nothing. Even if it's a Governance Token, who cares when the founders and 
VCs hold all the voting power? Anyways, I'm describing Airbnb Class A
Common Stock. Here's an excerpt from their S-1 [1] [2]:

> We have four series of common stock, Class A, Class B, Class C, and 
> Class H common stock (collectively, our "common stock"). The rights of 
> holders of Class A, Class B, Class C, and Class H common stock are 
> identical, except voting and conversion rights ... Each share of Class A 
> common stock is entitled to one vote, each share of Class B common stock 
> is entitled to 20 votes and is convertible at any time into one share of 
> Class A common stock ... Holders of our outstanding shares of Class B
> common stock will beneficially own 81.7% of our outstanding capital 
> stock and represent 99.0% of the voting power of our outstanding capital 
> stock immediately following this offering, ...

```
                   Name of             |  Class B   |   %   | % of Vot-
              Beneficial Owner         |   Shares   |       | ing Power
  -------------------------------------+------------+-------+-----------
    Brian Chesky                       | 76,407,686 | 29.1% |  27.1%
    Nathan Blecharczyk                 | 64,646,713 | 25.3% |  23.5%
    Joseph Gebbia                      | 58,023,452 | 22.9% |  21.4%
    Entities Affil. w/ Sequoia Capital | 51,505,045 | 20.3% |  18.9%     
```

Why do people buy tech stocks with inflated valuations? Some may because 
they believe that they will go up, that they will be more dominant, 
important, and valuable in the future. Like tokens, a large part of 
stocks' value is speculative. They are expressing their opinion on the 
future fundamentals. Others may simply because they believe others will 
believe that it is more valuable. Not fundamentals, this is an opinion 
about *pumpamentals*.

Importantly, unlike fundamental value, speculative value can be created 
out of thin air. It is minted by *fiat*. Fundamental value is difficult 
to create, whereas speculative value can be created through hype and 
psychology alone.

#### 3.3 - Shareholder Value

For stocks, there are usually laws in place to protect investors, pushing 
the balance between "speculation" and "fundamentals" towards the latter. 
As a result, firms are generally legally obligated to act in their 
shareholders' best interests. This is good because normal people will be 
able to participate in the wealth generated by companies. And obviously, 
companies should not defraud their investors.

However, the biggest *stake* holders in a business, are usually (in order):

1. The employees.  No matter what, no one else is spending 8 hours a day, 
   or ~33% of their total waking lifespan at this place. Whatever it is, 
   I guarantee you the employees feel it the most.

2. The customers.  The customers are the reason the business is able to 
   exist in the first place. Non-profits are not exempt: their customers 
   are their donors.

3. The local community / local environment / ecosystem.  The business
   doesn't exist in a vacuum. The business has externalities, and those 
   externalities affect most the immediate surrounding environment.

4. And in last place, the shareholders.  They do not really do anything 
   except contribute capital and hold the stock. Of course capital is 
   important but they are not spending 8 hours a day here, they are not 
   the reason the business exists, and in fact they might even live in a 
   totally different country.

For large, publicly-listed companies, the shareholders have one more 
unique difference from the other three stakeholders: liquidity. This 
difference is critical.

Liquidity describes how easy it is to buy and sell an asset. A dollar 
bill is liquid. Bitcoin is liquid. A house is relatively illiquid. Stock 
in large, publicly-listed companies is also liquid. A shareholder can buy 
a stock one day and sell it the next. As a result, the relationship is 
non-commital and opens the opportunity for short-term thinking. 

There are many things a company could do which would benefit shareholders 
short term, while harming the other three stakeholders long term. While a 
shareholder can simply dump their position and leave, the mess created is 
left for the employees, customers, and community to clean up.

(The SPAC boom was a pretty good example of this. Not all SPACs are bad, 
but a lot of pretty shit businesses publicly listed through SPACs then 
crashed. This is sad to me because some of that is early investors and 
founders dumping on retail like a crypto shitcoin, but dressed up because
it's NYSE or NASDAQ. Get liquidity then bail.)

Now, it is a misconception that stock companies must solely paperclip-
maximize short-term shareholder value. However, this is how it often 
plays out due to fucked up shit in the public markets, like annoying 
activist hedge funds or executive compensation tied to stock price. And 
it is true that employees can be shareholders. And that is usually a good 
thing! But few public companies are truly employee-owned.

Thinking about it from this perspective, the concept of maximizing 
shareholder value seems somewhat backwards. But *why* would one make 
this system where the priorities are seemingly inverted?

One benefit is that it would make your currency extremely valuable. 
Suppose you want to do some shit on Ethereum (speculating on some animal 
token?), you will need to have native ETH to do that transaction. 
Similarly, if you want to invest in US securities you at some point need
US Dollars. If you want to get a piece of that sweet $NVDA action, you
need dollars. People want to buy American stocks. American companies
perform well: they're innovative; they're not too heavily regulated;
it's a business friendly environment. (Shareholder value comes first!)
The numbers go up.

Remember the token founder from earlier in the Asian Arrangement? Suppose 
you are a *country* in the situation above, with a valuable currency. Not 
only is your currency in demand and valuable, you are the issuing/minting 
authority for that token. Similar to the token founder, you can print 
valuable money and pay for things with it.

And speaking of being a founder, let's talk about that!


### 4 - Startup Blues

Based on what we've set up so far, I will discuss some of the problems I 
see with many startups today and with startup culture.

Much of the problems stem from misalignment between shareholders and the 
other stakeholders (employees, etc). A lot of this comes from the 
fundamentals of venture capital. VC is itself an asset class, like fixed 
income and equities. VCs pitch this to their limited partners, at some 
level, based on the premise that their VC fund will generate yield for 
them. The strategy is to identify stuff that will become huge and buy it 
while it's still small and really cheap. Like trading shitcoins, it's
about finding what's going to moon and getting in early.

In a typical VC fund, a small handful of the investments will comprise the 
entire returns of the fund, with all of the other investments being 0's. 
The distribution is very power law. This means we are not looking for 1x, 
2x, or 3x outcomes; these may even be seen as failure modes. We are only 
interested in 20x, 50x, 100x, etc. outcomes. This is because anything 
less will be insufficient to make up for all the bad investments that 
get written down to zero.

For the same reason, it only makes sense for VCs to invest in certain 
types of companies. Have you ever heard this one? "We invest in SOFTWARE 
companies!...How is this SCALABLE? What do the VENTURE SCALE OUTCOMES look 
like here?" This is because these kinds of companies are the ones with the 
potential to 100x. They want you to deliver a 100x. Or how about this one? 
"We invest in CATEGORY-DEFINING companies". At least in security, 
"category-defining" means a shiny new checkbox in the compliance / cyber 
insurance questionnaire. In other words, a new kind of product that people 
MUST purchase. 

The market is incentivized to deliver a product that meets the minimum bar 
to meet that checkbox, while being useless. I invite you to think of your 
favorite middleware or EDR vendors here. For passionate security founders 
considering raising venture, remember that this is what your "success" is 
being benchmarked against.

```      
                      _.,------------------------------_ 
                   .%'                                 '&.  
                  .;'    We  partner  with  founders     ^;
                  !      building  category-defining      ;!
                  ;   companies at the earliest stages   _;
                   ^;                                  _.^
                     ''-.______________    __________.-' 
                                      /   /
                                     /  /^
                                    / /^
                                   /;^
                                  /' 
                   _________                           _________           
                _-'         '.                      _-'         '.         
              ,^             '^_                  ,^             '^_       
             /'               '"'                /'               '"'      
            ^'                 ^\^              ^'                 ^\^     
            :                   ^|              :                   ^|     
            :       .       .   |)              :       .       .   |)     
            :           \       |)              :           \       |)     
             :         __\     ,;                :         __\     ,;      
              "   !            ;                  "   !            ;       
              "   ^\  _____  /'                   "   ^\  _____  /'        
              '| | ^\      _/^                    '| | ^\      _/^         
               |    ^'====='                       |    ^'====='           
               | .   |   |                         | .   |   |             
             _'          |^__                    _'          |^__          
 ---------_-'        U       '--_ -------------_-'        U       '--_ -----  
 ._   _.-'                       '-._     _.-'                       '-    
   ':.'  \            ;         /     ': .'  \            ;         /    [4]
```

It's due to the thirst for 100x that there are painful dynamics. A 
fledgling startup may have founders they really like, but the current 
business may be unscalable. Bad VCs will push founders towards strategies, 
bets, models that have a 1% chance of working, but pay out 200x if they 
do. 

In the process they destroy a good business--one which has earned the 
trust of dutiful employees and loyal customers--all for a lottery ticket 
to build a unicorn. They will throw 100 darts at the dartboard and maybe 5 
will land, but what is it like to be the dart? You may have good expected 
value, but all of that EV is from spikes super far away from the origin. 
Is it pleasant betting everything on this distribution?

VC's want founders to be cult leaders. Have you ever heard this line? "We 
invest in great storytellers." Like what we saw with stocks and tokens, 
much of the easily-unlockable potential upside in assets is speculative. 
In essence, value can be created through narrative. Narrative *IS* value. 
Bad VC's will push founders to raise more capital at ever higher 
valuations (higher val = markup = fees), using narrative as fuel for the 
fire. Storytelling means "pump the token", and the job of the CEO is to 
(1) be the hype man and to raise (2) cash and (3) eyeballs. For this 
reason, Sam Altman and Elon are fine CEOs, regardless of other factors, 
because they are great at all three.

Much to the detriment of founders' and their employees' psyche, investors 
expect founders to be this legendary hype man. This requires a religiosity 
of belief that is borderline delusional. Have you ever tried to convince 
one of those Silicon Valley YC-type founder/CEOs that they are wrong? They 
will never listen to you because they have been socialized to be this way. 
It is what is expected of them, and it is easy to fall into this trap 
without even becoming aware of it. But if you think about it, does it make 
sense that to be a business owner, you need to be a religious leader? Of 
course not.

All of these reasons are why so many startup founders are young. They have 
little to lose, so gambling it all is OK. Being a cult leader may be 
traumatizing, but they have time (and the neuroplasticity) to heal. And 
lastly, they do not have the life experience to have a mature personal 
identity beyond "I am a startup founder". All of this makes it easy to 
accept the external pressures to build a company this or that way. And 
perhaps not the way they would have wanted to, relying instead on their 
personal values. The true irony is that the latter is what creates true, 
enduring company culture and not the made-up Mad Libs-tier Company Culture 
Notion Page shit that so many startups have. And of course, good VCs are 
self-aware of all of the issues and strive to prevent them. But the 
overall problem remains.

One last externality is for communities based around an industry. When you 
add billions of venture dollars into an industry, it becomes cringe. 
It's saddening to me seeing the state of certain cybersecurity conferences 
which are now dominated by..."COME TO OUR BOOTH, YOU CAN BE A HACKER. 
PLEASE VIEW OUR AI GENERATED GRAPHICS OF FIGURES CLAD IN DARK HOODIES 
STATIONED BEHIND LAPTOPS". Here I would use the pensive emoji U+1F614 
to describe my feelings about the appropriation of hacker culture but 
Phrack is 7-bit ASCII, so please have this: :c u_u . _.

### 5 - Takeaways

The point is, all of this made me feel very small and powerless after I 
realized the sheer size of the problems I was staring at. Nowadays, to 
me it's about creating good jobs for my friends, helping our customers, 
and taking care of the community. Importantly, I realized that this is 
still making a bigger positive impact than what I could have done alone 
just as an individual hacker or engineer.

To me, businesses are economic machines that can create positive (or 
negative) impact in a consistent, self-sustaining way. There are many 
people who are talented, kind, and thoughtful but temporarily unlucky. 
Having a company let me help these friends monetize their abilities and be 
rewarded fairly for them. And in that way I helped make their life better. 
Despite a lot of the BS involved in running a business, this is one thing 
that is very meaningful to me.

You can understand computers and science and math as much as you want, but 
you will not be able to fix the bigger issues by yourself. The systems 
that run the world are much bigger than what we can break on our laptops 
and lab benches.

But like those familiar systems, if we want to change things for the 
better, we have to first understand those systems. Knowledge is power. 
Understanding is the first step towards change. If you do not like the 
system as it is, then it is your duty to help fix it.

Do not swallow blackpills. It's easy to get really cynical and think 
things are doomed (to AGI apocalypse, to environmental disaster, to 
techno/autocratic dystopia, whatever). I want to see a world where 
thoughtful hackers learn these systems and teach each other about them. 
That generation of hackers will wield that apparatus, NOT THE OTHER WAY 
AROUND.

Creating leverage for yourself.  Hackers should not think of themselves as 
"oh I am this little guy fighting Big Corporation" or whatever. This is 
low agency behavior. Instead become the corporation and RUN IT THE WAY YOU 
THINK IT SHOULD BE RUN. Keep it private and closely held, so no one can 
fuck it up. Closely train up successors, so in your absence it will 
continue to be run in a highly principled way that is aligned with your 
values and morals. Give employees ownership, as it makes everyone aligned 
with the machine's long-term success, not just you.

Raising capital.  Many things do really need capital, but raise in a 
responsible way that leaves you breathing room and the freedom to operate 
in ways that are aligned with your values. Never compromise your values or 
integrity. Stay laser focused on cash flows and sustainability, as these 
grant you the freedom to do the things right.

HACKERS SHOULDN'T BE AFRAID TO TOUCH THE CAPITAL MARKETS.  Many hackers 
assume "oh that fundraising stuff is for charismatic business types". I 
disagree. It's probably better for the world if good thoughtful hackers 
raise capital. Giving them leverage to change the world is better than 
giving that leverage to some psycho founder drinking the Kool-Aid. I 
deeply respect many of the authors in Phrack 71, and I would trust them to 
do a better job taking care of things than an amorphous amalgam of angry 
and greedy shareholders.

For all things that don't need capital, do not raise. Stay bootstrapped 
for as long as possible. REMEMBER THAT VALUATION IS A VANITY METRIC. Moxie 
Marlinspike wrote on his blog [3] that we are often guilty of always
trying to quantify success. But what is success? You can quantify net
worth, but can you quantify the good you have brought to others lives?

For personal goals, think long term. People tend to overestimate what they 
can do in 1 year, but underestimate what they can do in 10. DO NOT start a 
company thinking you can get your hands clean of it in 2-3 years. If you 
do a good job, you will be stuck with it for 5-10+ years. Therefore, DO 
NOT start a company until you are sure that is what you want to do with 
your life, or at least, your twenties/thirties (depending on when you 
start). A common lament among founders, even successful ones, is: 
"Sometimes I feel like I'm wasting my twenties". There's an easy Catch-22 
here: you may not know what you really want until you do the company; but 
once you do the company, you won't really be able to get out of it. Be 
wary of that.

Creating value.  This is one of those meaningless phrases that I dislike. 
Value is what you define it to be. Remember to work on things that have 
TAMs, but remember that working on art is valuable too! It is not all 
about the TAM monster--doing cool things that are NOT ECONOMICALLY 
VALUABLE, but ARTISTICALLY VALUABLE, is equally important. There is not 
much economic value in a beautiful polyglot file, but it is artistically 
delightful. This is part of why people hate AI art: it may be economically 
valuable, but it is often artistically bankrupt. (Some people do use 
generative tools in actually original and artistic ways, but this is the 
exception not the norm currently.)

Founders vs Investors.  Here is my advice: Ignore any pressure from 
investors to make company "scalable" or whatever. Make sure your investors 
have no ability to fire you or your co-founder(s). Make sure you and 
co-founder are always solid and trust each other more than investors. You 
and your cofounders need to be BLOOD BROTHERS (/sisters/w.e). If an 
investor is trying to play politics with one of you to go against the 
other cofounder, cut that investor out immediately and stop listening to 
them.

Any investor who pushes for scalability over what you think is the best 
interest of the company is not aligned with you. High-quality investors 
will not push for this because they are patient and in it for the long 
game. If you are patient, you can make a very successful company, even if
it is not that scalable. High-quality investors will bet on founders and
are committed; only bad ones will push for this kind of shit.

I'm going to avoid giving more generic startup advice here. Go read Paul 
Graham's essays. But remember that any investor's perspective will not be
the perspective of you and your employees. Pivoting 5 times in 24 months
is not a fun experience to work at: your employees will resign while your
investors celebrate your "coming of age journey"--unless everyone signed
up for that terrifying emotional rollercoaster from the start.

They say that "hacker" is a dying identity. Co-opted by annoying VC-backed 
cybersecurity companies that culturally appropriate the identity, the term 
is getting more polluted and diluted by the day. Meanwhile, computers are 
getting more secure, and they are rewriting everything in Rust with 
pointers-as-capability machines and memory tagging. Is it over?

I disagree. As long as the hacker *ethos* is alive, regardless of any 
particular scene, the identity will always exist. However, now is a 
crucible moment as a diaspora of hackers, young and old, venture out into 
the world. 

Calling all hackers: never forget who you are, who you will become, and 
the mark you leave.


### 6 - Thanks

Greetz (in no particular order):
 * ret2jazzy, Sirenfal, ajvpot, rose4096, Transfer Learning, samczsun, 
   tjr, claire (aka sport), and psifertex.
 * perfect blue, Blue Water, DiceGang, Shellphish, and all CTF players.
 * NotJan, nspace, xenocidewiki, and the members of pinkchan and Secret Club.
 * Everyone at Zellic, past and present.

Finally, a big thank you to the Phrack staff (shoutout to netspooky and 
richinseattle!) for making this all possible.


### 7 - References

[1] https://www.sec.gov/Archives/edgar/data/1559720/000119312520315318/
        d81668d424b4.htm
[2] https://www.sec.gov/Archives/edgar/data/1559720/000119312522115317/
        d278253ddef14a.htm
[3] https://moxie.org/stories/promise-defeat/

[4] https://twitter.com/nikitabier/status/1622477273294336000


### 8 - Appendix: Financial institution glossary for hackers

(Not serious! For jokes... :-)

- IB:  Investment Bank. Basically collect fat fees to do up ("advise on") 
       M&As and other transactions. Help match buyers and sellers for your 
       private equity. They are like CYA for your deal.

- PE:  Private Equity. Basically buy not-overly-seriously ("poorly") run 
       companies, fire the management, then run it "professionally" (i.e. 
       make it generally shitty for customers and employees and community 
       for the benefit of shareholders)

- HF:  Hedge Fund. Trade out pricing inefficiencies

- MM:  Market Maker. Basically the same thing

- VC:  Basically gamble on tokens (crypto or stocks) and back cool and/or 
       wacky ideas that the rest of these people find too stinky to invest
       in

- PnD: Pump and Dump.

- TVL: Total Value Locked. Basically how much money is currently in a 
       blockchain or smart contract system.

- TPS: Transactions Per Second. A measure of how scalable or useful a 
       blockchain or database is. An oft-abused metric hacked by vaporware 
       shillers for hype and PnD purposes.

- TAM: Total Addressable ~~Memory~~ Market. Basically how much money a 
       given idea can make.

- NFA: Not finanical advice.

|=[ EOF ]=---------------------------------------------------------------=|
