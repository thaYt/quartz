const COLOR_BLACK = "#000000";
const COLOR_DARK_BLUE = "#0000AA";
const COLOR_DARK_GREEN = "#00AA00";
const COLOR_DARK_AQUA = "#00AAAA";
const COLOR_DARK_RED = "#AA0000";
const COLOR_DARK_PURPLE = "#AA00AA";
const COLOR_GOLD = "#FFAA00";
const COLOR_GRAY = "#AAAAAA";
const COLOR_DARK_GRAY = "#555555";
const COLOR_BLUE = "#5555FF";
const COLOR_GREEN = "#55FF55";
const COLOR_AQUA = "#55FFFF";
const COLOR_RED = "#FF5555";
const COLOR_LIGHT_PURPLE = "#FF55FF";
const COLOR_YELLOW = "#FFFF55";
const COLOR_WHITE = "#FFFFFF";

function getLevelColor(level) {
  if (level >= 1000) {
    return COLOR_DARK_GRAY;
  } else if (level >= 900) {
    return COLOR_DARK_PURPLE;
  } else if (level >= 800) {
    return COLOR_BLUE;
  } else if (level >= 700) {
    return COLOR_LIGHT_PURPLE;
  } else if (level >= 600) {
    return COLOR_DARK_RED;
  } else if (level >= 500) {
    return COLOR_DARK_AQUA;
  } else if (level >= 400) {
    return COLOR_DARK_GREEN;
  } else if (level >= 300) {
    return COLOR_AQUA;
  } else if (level >= 200) {
    return COLOR_GOLD;
  } else if (level >= 100) {
    return COLOR_WHITE;
  } else {
    return COLOR_GRAY;
  }
}

function getFKDRColor(fkdr) {
  if (fkdr >= 10) {
    return COLOR_DARK_RED;
  } else if (fkdr >= 5) {
    return COLOR_RED;
  } else if (fkdr >= 3) {
    return COLOR_GOLD;
  } else if (fkdr >= 2) {
    return COLOR_YELLOW;
  } else if (fkdr >= 1) {
    return COLOR_GREEN;
  } else {
    return COLOR_DARK_GREEN;
  }
}

function getWLRColor(wlr) {
  if (wlr >= 5) {
    return COLOR_DARK_RED;
  } else if (wlr >= 2.5) {
    return COLOR_RED;
  } else if (wlr >= 1.5) {
    return COLOR_GOLD;
  } else if (wlr >= 1) {
    return COLOR_YELLOW;
  } else if (wlr >= 0.5) {
    return COLOR_GREEN;
  } else {
    return COLOR_DARK_GREEN;
  }
}

function getBBLRColor(bblr) {
  if (bblr >= 5) {
    return COLOR_DARK_RED;
  } else if (bblr >= 2.5) {
    return COLOR_RED;
  } else if (bblr >= 1.5) {
    return COLOR_GOLD;
  } else if (bblr >= 1) {
    return COLOR_YELLOW;
  } else if (bblr >= 0.5) {
    return COLOR_GREEN;
  } else {
    return COLOR_DARK_GREEN;
  }
}

function getFinalKillsColor(finalKills) {
  if (finalKills >= 1000) {
    return COLOR_WHITE;
  } else {
    return COLOR_GRAY;
  }
}

function getWinsColor(wins) {
  if (wins >= 100) {
    return COLOR_WHITE;
  } else {
    return COLOR_GRAY;
  }
}

function getIndexColor(index) {
  if (index == "fold rn") {
    return COLOR_BLACK;
  }
  if (index >= 10000) {
    return COLOR_DARK_RED;
  } else if (index >= 5000) {
    return COLOR_RED;
  } else if (index >= 2500) {
    return COLOR_GOLD;
  } else if (index >= 1500) {
    return COLOR_YELLOW;
  } else if (index >= 750) {
    return COLOR_GREEN;
  } else {
    return COLOR_DARK_GREEN;
  }
}

function getRankColor(rank) {
  console.log(rank);
  switch (rank) {
    case "NON":
      return COLOR_GRAY;
    case "VIP":
    case "VIP+":
      return COLOR_GREEN;
    case "MVP":
    case "MVP+":
      return COLOR_AQUA;
    case "MVP++":
      return COLOR_GOLD;
    case "GAME_MASTER":
      return COLOR_DARK_GREEN;
    case "ADMIN":
      return COLOR_DARK_RED;
    case "YOUTUBE":
      return COLOR_RED;
  }
}

export function getColor(stat: string, data: any): string {
  console.log(stat, data)
  switch (stat) {
    case "Name":
      return getRankColor(data);
    case "Finals":
      return getFinalKillsColor(data);
    case "Wins":
      return getWinsColor(data);
    case "Losses":
      return getWinsColor(data);
    case "Beds":
      return getFinalKillsColor(data);
    case "Bl":
      return getFinalKillsColor(data);
    case "FKDR":
      return getFKDRColor(data);
    case "WLR":
      return getWLRColor(data);
    case "Level":
      return getLevelColor(data);
    case "BBLR":
      return getBBLRColor(data);
    default:
      return "white";
  }
}
