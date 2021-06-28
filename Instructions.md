![manual](E:\programming\go\projects\src\about_kid_stealing_back\manual.png)



**CONTROLS**

The default control scheme:

```
Move up:    up arrow
Move down:  down arrow
Move left:  left arrow
Move right: right arrow
Wait:       space

Look:    l
Pick up: g
drop:    d
throw:   f

contextual environment action: Enter
```

You can customize the control scheme by the changing `options_controls.cfg` file. There, you will need to set `CUSTOM_CONTROLS` to `TRUE`, and then edit the control scheme below.

This game also provides (rather experimental and not very well tested) support for multiple keyboard layouts: QWERTY (the default one), QWERTZ, AZERTY, DVORAK.



**GAMEPLAY**

Your goal it to stole (or maybe *recover*?) as many treasures as possible, and stay alive. You can not fight back, so you need to be clever: hide under the tables, listen to surroundings, and throw pebbles at the enemies to buy some time if necessary. The stolen valuables you can store in the hatch to the old tunnel – you need to enter it (to do it, press the key that triggers contextual environmental action).



**UI**

![image-20210313004631899](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313004631899.png) – it is your HP; if it drops to 0, you are dead.

![image-20210313004914524](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313004914524.png) – indicated amount of pebbles in your pockets; you can stagger Vikings by throwing pebbles at them.

![image-20210313005202819](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313005202819.png) – the yellow icons represent treasures in your current possession; the gray ones symbolize free inventory slots.

![image-20210313005322396](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313005322396.png) – number of carried items (excluding pebbles) affect encumbrance.



**MAP**

![image-20210313005454076](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313005454076.png) – impassable wall.

![image-20210313005552987](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313005552987.png) – floor.

![image-20210313005735828](E:\programming\go\projects\src\about_kid_stealing_back\screenshots\image-20210313005735828.png) – you can hide under the table.

![image-20210313005822012](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313005822012.png) – chairs do not allow hidings.

![image-20210313005919315](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313005919315.png) – you can hide and sneak through the crumbling walls; the enemies can not follow you there.

![image-20210313010048667](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010048667.png) – pillar; simple obstacle.

![image-20210313010630915](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010630915.png) – hatch to the old tunnel; it is a safe place to hide all the treasures.

![image-20210313010123651](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010123651.png) – pebbles; you can pick it up, or throw at the enemies directly from the floor.

![image-20210313010317043](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010317043.png) – enemy unaware of your presence.

![image-20210313010428740](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010428740.png) – enemy that is actually chasing you.

![image-20210313010513859](C:\Users\Ved\AppData\Roaming\Typora\typora-user-images\image-20210313010513859.png) – staggered enemy.



**POINTS**

This game features High Scores table. While the exact amount of points added or deducted by specific activities will not be written here, below there is info about all factors provided.

Points are added when:

- player stores the relics in the hatch,
- extra points for stealing all the valuables.

Points are deducted when:

- player die,
- player ends the game without any valuable stolen.



**HINTS**

- Hide (inside the crumbling wall, under the tables) from the  enemies; a hidden player is dark gray, almost black, and the player icon changes; to hide  successfully, the player has to enter the hiding tile when there are no  triggered enemies in the field of view; otherwise, player will not be  hidden, and the color of player character will be a lighter gray.
- If discovered, run - as a kid, you can actually squeeze through crumbling walls, and the Vikings can not follow the same path.
- To lose the enemies, throw pebbles at them from your own pocket or directly from the ground.
- If your bag is full, enter the hatch, store the valuables there, and go  back to the cellar; you can not steal all the treasures at once, but it  is guaranteed that clearing the level is possible in two runs is  possible.
- Mind your encumbrance – with each container filled, you are becoming slower; still, the bigger risk, the bigger rewards – the  scoring system promotes clearing level in as small number of attempts as possible.