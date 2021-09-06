article: https://github.com/refinedcoding/CruelSystem/blob/main/homework/2021-08/0816/whatsapp%20multi-device.md

situation:   phone is the source of truth for all user data and the only device capable of end-to-end encrypting messages for another user, initiating calls, etc. Companion devices maintain a persistent secure connection with the phone and simply mirror its contents on their own UI.

tradeoff:

    1 reliability trade-offs: companion devices are slower and sometimes get disconnected in every operation, since they needs the phone to actually operates everything. Imaging if someone phone has poor connection but have a really fast ethernet on the computer, it will be really slow and annoying.

    2 security tradeoff: It also allows for only a single companion device to be operative at a time. 

actions to maintain security: 

1each person has an identity key -> each person-device has an identity key

2 We have also addressed the challenge of preventing a malicious or compromised server from eavesdropping on someone’s communications by surreptitiously adding devices to someone’s account.

2.1 security codes to represent all of one's device identities ? how does it prevents linking a malacious device? 

2.2 Automatic Device Verification to establish trust between each other, security code only needed when adding a totally new device (how?)

2.3 add biometrics requirement before scanning QR code to add device (ok) 

2.4 user can see their own device list and remote kick out a device (ok)



不看不知道，一看吓一跳，原来whatsapp上有2billion 用户，比微信的用户数还多。。

![image-20210816201835703](https://tva1.sinaimg.cn/large/008i3skNgy1gtjm9vadd0j610t0u0abw02.jpg)
