<script lang="ts">
    let {videoSrc}:{videoSrc:string}=$props();
  </script>
  
  <video width="640" height="360" controls>
    <source src={videoSrc} type="video/mp4">
    Ваш браузер не поддерживает видео.
  </video>
  