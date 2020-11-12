<template>
    <div id="noVNC_all" class="vnc-card">
        <div id="noVNC_status_bar">
            <div id="noVNC_left_dummy_elem"></div>
<!--            <div id="noVNC_status">Loading: {{url}}</div>-->
<!--            <div id="noVNC_buttons">-->
<!--                <input type=button value="Send CtrlAltDel"-->
<!--                       id="sendCtrlAltDelButton" class="noVNC_shown">-->
<!--                <span id="noVNC_power_buttons" class="noVNC_hidden">-->
<!--        <input type=button value="Shutdown"-->
<!--               id="machineShutdownButton">-->
<!--        <input type=button value="Reboot"-->
<!--               id="machineRebootButton">-->
<!--        <input type=button value="Reset"-->
<!--               id="machineResetButton">-->
<!--      </span>-->
            </div>
        </div>
    </div>
</template>

<script>

import RFB from '@novnc/novnc/core/rfb'
export default {
  name: 'novnc',
  props: ['Vnc', 'viewOnly'],
  data () {
    return {
      rfb: null,
      IsClean: false, // 是否已断开并不可重新连接
      connectNum: 0 // 重连次数
    }
  },
  computed: {
    url: function () {
      // `this` указывает на экземпляр vm
      return `ws://172.31.50.74:4444/vnc/${this.Vnc}`
    }
  },
  methods: {
    // vnc连接断开的回调函数
    disconnectedFromServer (msg) {
      console.log('disconnectedFromServer')
      if (msg.detail.clean) {
        // 根据 断开信息的msg.detail.clean 来判断是否可以重新连接
        this.contentVnc()
      } else {
        // 这里做不可重新连接的一些操作
      }
    },
    // 连接成功的回调函数
    connectedToServer () {
      console.log('success')
    },

    // 连接vnc的函数
    connectVnc () {
      const PASSWORD = 'selenoid'
      let rfb = new RFB(document.querySelector('#noVNC_all'), this.url, {
        // 向vnc 传递的一些参数，比如说虚拟机的开机密码等
        credentials: { password: PASSWORD }
      })
      rfb.addEventListener('connect', this.connectedToServer)
      rfb.addEventListener('disconnect', this.disconnectedFromServer)
      rfb.scaleViewport = true // scaleViewport指示是否应在本地扩展远程会话以使其适合其容器。禁用时，如果远程会话小于其容器，则它将居中，或者根据clipViewport它是否更大来处理。默认情况下禁用。
      rfb.resizeSession = true // 是一个boolean指示是否每当容器改变尺寸应被发送到调整远程会话的请求。默认情况下禁用
      rfb.viewOnly = this.viewOnly;
      this.rfb = rfb
    }
  },
  // 在mounted周期里面连接vnc
  mounted () {
    // 这时已经可以获取到dom元素
    console.log('在mounted周期里面连接vnc')
    this.connectVnc()
  }

}
</script>

<style scoped>
    .vnc-card {
      height: 500px;
        width: 100%;
        display: flex;
        flex-direction: column;
        box-shadow: 0 1px 6px rgba(0, 0, 0, 0.12), 0 1px 4px rgba(0, 0, 0, 0.12);
    }
</style>
