<script>
  export const durationUnitRegex = /[a-zA-Z]/;

  export const calculateRgba = (color, opacity) => {
    if (color[0] === "#") {
      color = color.slice(1);
    }

    if (color.length === 3) {
      let res = "";
      color.split("").forEach((c) => {
        res += c;
        res += c;
      });
      color = res;
    }

    const rgbValues = (color.match(/.{2}/g) || [])
      .map((hex) => parseInt(hex, 16))
      .join(", ");

    return `rgba(${rgbValues}, ${opacity})`;
  };

  export const range = (size, startAt = 0) =>
    [...Array(size).keys()].map(i => i + startAt);

  export let color = "#FF3E00";
  export let unit = "px";
  export let duration = "2.1s";
  export let size = "60";
  let rgba;
  $: rgba = calculateRgba(color, 0.2);
</script>

<style>
  .wrapper {
    height: calc(var(--size) / 15);
    width: calc(var(--size) * 2);
    background-color: var(--rgba);
    position: relative;
    overflow: hidden;
    background-clip: padding-box;
  }
  .lines {
    height: calc(var(--size) / 15);
    background-color: var(--color);
  }
  .small-lines {
    position: absolute;
    overflow: hidden;
    background-clip: padding-box;
    display: block;
    border-radius: 2px;
    will-change: left, right;
    animation-fill-mode: forwards;
  }
  .small-lines.\31 {
    animation: var(--duration) cubic-bezier(0.65, 0.815, 0.735, 0.395) 0s
      infinite normal none running long;
  }
  .small-lines.\32 {
    animation: var(--duration) cubic-bezier(0.165, 0.84, 0.44, 1)
      calc((var(--duration)+0.1) / 2) infinite normal none running short;
  }
  @keyframes long {
    0% {
      left: -35%;
      right: 100%;
    }
    60% {
      left: 100%;
      right: -90%;
    }
    100% {
      left: 100%;
      right: -90%;
    }
  }
  @keyframes short {
    0% {
      left: -200%;
      right: 100%;
    }
    60% {
      left: 107%;
      right: -8%;
    }
    100% {
      left: 107%;
      right: -8%;
    }
  }
</style>

<div class="wrapper" style="--size: {size}{unit}; --rgba:{rgba}">
  {#each range(2, 1) as version}
    <div
      class="lines small-lines {version}"
      style="--color: {color}; --duration: {duration};" />
  {/each}
</div>