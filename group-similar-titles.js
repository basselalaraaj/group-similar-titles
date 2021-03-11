const alphabet = "abcdefghijklmnopqrstuvwxyz";

const groupTitles = (titles) => {
  const groupedTitles = new Map();
  titles.forEach((title) => {
    const wordMap = new Array(26).fill(0);
    [...title].forEach((letter) => {
      const letterIndex = alphabet.indexOf(letter);
      if (letterIndex >= 0) {
        wordMap[letterIndex]++;
      }
    });

    const titleKey = "#" + wordMap.join("#");
    const titlesByKey = groupedTitles.get(titleKey);
    const newTitles = titlesByKey ? [...titlesByKey, title] : [title];

    groupedTitles.set(titleKey, newTitles);
  });

  return groupedTitles.values();
};

const titles = ["duel", "dule", "speed", "spede", "deul", "cars"];
const query = "spede";
for (const groupedTitles of groupTitles(titles)) {
  if (groupedTitles.includes(query)) {
    console.log(groupedTitles);
  }
}
