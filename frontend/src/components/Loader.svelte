<script>
  export const durationUnitRegex = /[a-zA-Z]/;
  export const range = (size, startAt = 0) =>
  [...Array(size).keys()].map(i => i + startAt);
  export let color = "#ffffff";
  export let unit = "px";
  export let duration = "1.5s";
  export let size = "60";
  let durationUnit = duration.match(durationUnitRegex)[0];
  let durationNum = duration.replace(durationUnitRegex, "");
</script>

<style>
  .wrapper {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    width: var(--size);
    height: calc(var(--size) / 2.5);
    margin-top: calc(var(--size) / 5);
  }
  .circle {
    position: absolute;
    top: 0px;
    width: calc(var(--size) / 5);
    height: calc(var(--size) / 5);
    border-radius: 999px;
    background-color: var(--color);
    animation: motion var(--duration) cubic-bezier(0.895, 0.03, 0.685, 0.22)
      infinite;
  }
  @keyframes motion {
    0% {
      opacity: 1;
    }
    50% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }
</style>

<div
  class="wrapper"
  style="--size: {size}{unit}; --color: {color}; --duration: {duration}">
  {#each range(3, 0) as version}
    <div
      class="circle"
      style="animation-delay: {version * (+durationNum / 10)}{durationUnit}; left: {version * (+size / 3 + +size / 15) + unit};" />
  {/each}
</div>
