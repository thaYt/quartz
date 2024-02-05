<script lang="ts">
  import type { BedwarsPlayer } from "src/types";
  import { EventsOn } from "../wailsjs/runtime";
  import Player from "./Player.svelte";
  import { onMount } from "svelte";

  let players: BedwarsPlayer[] = [];
  $: players = players.sort((a, b) => b.Level - a.Level); // todo customizable sorting

  let layout = [ // todo customizable layout
    "Level",
    "Name",
    "FKDR",
    "Finals",
    "Wins",
    "Losses",
    "Beds",
    "Bl",
    "WLR",
  ];

  let customWidths = { // todo customizable widths
    Name: "2fr",
    Bl: ".5fr",
  };

  let gridTemplateColumns = layout
    .map((prop) => customWidths[prop] || "1fr")
    .join(" ");

  onMount(() => {
    EventsOn("addPlayer", (data: BedwarsPlayer) => {
      console.log(data);
      players = [...players, data];
    });

    EventsOn("removePlayer", (data: BedwarsPlayer) => {
      players = players.filter((player) => player.Name !== data.Name);
    });

    EventsOn("nuke", () => {
      players = [];
    });
  });
</script>

<div
  class="grid-container"
  style="grid-template-columns: {gridTemplateColumns};"
>
  <div class="header">
    {#each layout as option (option)}
      <div class="headerItem">{option}</div>
    {/each}
  </div>
  {#each players as player (player.Name)}
    <Player {player} {layout} />
  {/each}
</div>

<style lang="scss">
  .header {
    display: contents;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    height: 20px;
    line-height: 32px;
    font-size: 16px;
  }

  .headerItem {
    grid-column: span 1;
    text-align: center;
    background-color: rgba(15, 15, 15, 0.2);
  }

  .grid-container {
    display: grid;
    width: 100%;
  }
</style>
