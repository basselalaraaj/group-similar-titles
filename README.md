# group-similar-titles
If a user misspells a word for example `spede` instead of `speed`, they will still be shown the correct title. Searching for `spede` will return both `speed` and `spede`.

We compute a 26-element vector for each English letter. This mapping process generates identical vectors for the strings that are anagrams. We represent `speed` or `spede` as `#0#0#0#1#2#0#0#0#0#0#0#0#0#0#0#1#0#0#1#0#0#0#0#0#0#0`. We then search the map and return all the map entries using the vector.