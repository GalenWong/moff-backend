# MOFF (backend)

MOFF stands for Music offline. 
I aim to create a music app platform that would allow me to sync the music file on my laptop with my phone.

I travel a lot on airplane. 
Not having internet access means I cannot listen to songs conveniently on my Phone/Laptop.
Then, I have to download music. Usually, it is done on my laptop.
What often annoys me more is that the songs I have on my laptop is not synchronised with the songs on my phone. 

Since iOS is quite a "enclosed" system, importing music from laptop into iPhone is not a good experience IMO. 

Therefore, I created MOFF to combat this situation.

## The Stack
This repo is for the backend component of MOFF. 
I wrote it in Golang but not JavaScript since JS is just a bit too complicated for me.
I use Mongo as the database. 
To store the song on the Cloud so I can sync the files, I used Google Drive as the main storage.
