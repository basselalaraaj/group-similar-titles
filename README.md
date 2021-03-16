# group-similar-titles
If a user misspells a word for example `spede` instead of `speed`, they will still be shown the correct title. Searching for `spede` will return both `speed` and `spede`.

We compute a 26-element vector for each English letter. This mapping process generates identical vectors for the strings that are anagrams. We represent `speed` or `spede` as `#0#0#0#1#2#0#0#0#0#0#0#0#0#0#0#1#0#0#1#0#0#0#0#0#0#0`. We then search the map and return all the map entries using the vector.

<img width="584" alt="Schermafbeelding 2021-03-16 om 22 33 31" src="https://user-images.githubusercontent.com/5745279/111382758-a91f6700-86a7-11eb-9b57-3c121025eccf.png">