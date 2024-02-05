<script lang="ts">
  import type { BedwarsPlayer } from "src/types";

  import { getColor } from "./assets/color";

  export let player: BedwarsPlayer;
  export let layout: string[];
</script>

<div class="player">
  {#each layout as property}
    {#if property === "Name"}
      <div class="property" style="color: {getColor(property, player.Rank)};">
        {player[property]}
      </div>
    {:else}
      <div
        class="property"
        style="color: {getColor(property, player[property])};"
      >
        {player[property]}
      </div>
    {/if}
  {/each}
</div>

<style lang="scss">
  .player {
    display: contents;
  }

  .property {
    display: flex;
    align-items: center;
    justify-content: center;
  }
</style>
