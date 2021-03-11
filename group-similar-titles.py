def groupTitles(titles):
    groupedTitles = {}
    for title in titles:
        wordMap = [0]*26
        for c in title:
            index = ord(c) - ord("a")
            wordMap[index] += 1
        titleKey = "#".join(map(str, wordMap))
        if titleKey in groupedTitles:
            groupedTitles[titleKey].append(title)
        else:
            groupedTitles[titleKey] = [title]
    return groupedTitles.values()


titles = ["duel", "dule", "speed", "spede", "deul", "cars"]
query = "spede"

for groupedTitles in groupTitles(titles):
    if(query in groupedTitles):
        print(groupedTitles)
