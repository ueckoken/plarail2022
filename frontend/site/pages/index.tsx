import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import RailroadMap from "../components/RailRoadMap"
import VideoCast from "../components/VideoCast"
import { useEffect, useRef, useState } from "react"
import {
  BlocklId,
  bunkiRailId,
  BunkiRailId,
  Message,
  StationId,
  StationState,
  StopRailId,
  stopRailId,
} from "../types/control-messages"
import SpeedMeter from "../components/svgParts/SpeedMeter"
import { SpeedMessage, TrainId } from "../types/speed-messages"
import ReverseHandle from "../components/svgParts/ReverseHandle"

// OFF: false, ON: trueと対応
type StopPointState = Record<StopRailId, boolean>
const INITIAL_STOP_POINT_STATE: StopPointState = {
  motoyawata_s1: false,
  motoyawata_s2: false,
  iwamotocho_s1: false,
  iwamotocho_s2: false,
  iwamotocho_s4: false,
  kudanshita_s5: false,
  kudanshita_s6: false,
  sasazuka_s1: false,
  sasazuka_s2: false,
  sasazuka_s3: false,
  sasazuka_s4: false,
  sasazuka_s5: false,
  meidaimae_s1: false,
  meidaimae_s2: false,
  chofu_s1: false,
  chofu_s2: false,
  chofu_s3: false,
  chofu_s4: false,
  chofu_s5: false,
  chofu_s6: false,
  kitano_s1: false,
  kitano_s2: false,
  kitano_s3: false,
  kitano_s4: false,
  kitano_s5: false,
  kitano_s6: false,
  kitano_s7: false,
  takao_s1: false,
  takao_s2: false,
}

type BlockState = Record<BlocklId, boolean>
const INITIAL_BLOCK_STATE: BlockState = {
  shinjuku_b1: false,
  shinjuku_b2: false,
  sakurajosui_b1: false,
  sakurajosui_b2: false,
  sakurajosui_b3: false,
  sakurajosui_b4: false,
  sakurajosui_b5: false,
  sakurajosui_b6: false,
  chofu_b1: false,
  chofu_b2: false,
  chofu_b3: false,
  chofu_b4: false,
  chofu_b5: false,
  hashimoto_b1: false,
  hashimoto_b2: false,
  hachioji_b1: false,
  hachioji_b2: false,
}

type SwitchPointState = Record<BunkiRailId, boolean>
const INITIAL_SWITCH_POINT_STATE: SwitchPointState = {
  chofu_p1: false,
  sakurajosui_p1: false,
}

const Home: NextPage = () => {
  const stationWs = useRef<WebSocket>()
  const [stopPointState, setStopPointState] = useState<StopPointState>(
    INITIAL_STOP_POINT_STATE
  )
  const [switchPointState, setSwitchPointState] = useState<SwitchPointState>(
    INITIAL_SWITCH_POINT_STATE
  )
  useEffect(() => {
    setSwitchPointState(INITIAL_SWITCH_POINT_STATE)
  })
  const [blockState, setBlockState] = useState<BlockState>(INITIAL_BLOCK_STATE)
  const [selectedStationId, setSelectedStationId] =
    useState<StationId>("unknown")
  const [trainPosition1, setTrainPosition1] = useState<number>(0.4)

  const speedWs = useRef<WebSocket>()

  const [isBack, setIsBack] = useState<boolean>(false)

  const [roomIds, setRoomIds] = useState<string[]>(["chofu", "train"])

  const changeStopPointOrSwtichPointState = (
    stationId: StationId,
    state: StationState
  ) => {
    const message: Message = {
      station_name: stationId,
      state: state,
    }
    stationWs.current?.send(JSON.stringify(message))
  }
  const toggleStopPointOrSwitchPointState = (stationId: StationId) => {
    let state
    if (stopRailId.is(stationId)) {
      state = stopPointState[stationId]
    } else if (bunkiRailId.is(stationId)) {
      state = switchPointState[stationId]
    } else {
      return
    }
    const nextState = state ? "OFF" : "ON"
    changeStopPointOrSwtichPointState(stationId, nextState)
  }

  useEffect(() => {
    const ws = new WebSocket("wss://speed.chofufes2022.ueckoken.club/speed")
    speedWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
    })
    ws.addEventListener("message", (e) => {
      const message: SpeedMessage = JSON.parse(e.data)
      console.log(message)
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const ws = new WebSocket("wss://control.chofufes2022.ueckoken.club/ws")
    stationWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
      console.log(e)
    })
    ws.addEventListener("message", (e) => {
      console.log("recieved message")
      console.log(e)
      const message: Message = JSON.parse(e.data)
      console.log(message)
      if (message.station_name === "unknown" || message.state === "UNKNOWN") {
        return
      }
      if (stopRailId.is(message.station_name)) {
        setStopPointState((previousStopPointState) => ({
          ...previousStopPointState,
          [message.station_name]: message.state === "ON",
        }))
        return
      }
      if (bunkiRailId.is(message.station_name)) {
        setSwitchPointState((previousSwitchPointState) => ({
          ...previousSwitchPointState,
          [message.station_name]: message.state === "ON",
        }))
      }
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const intervalId = setInterval(() => {
      const tmpTrainPosition1 = trainPosition1 + 0.01
      setTrainPosition1(tmpTrainPosition1 <= 1 ? tmpTrainPosition1 : 0)
    }, 20)
    return () => clearInterval(intervalId)
  })

  return (
    <div className={styles.container}>
      <Head>
        <title>工研&times;鉄研プラレール展示 操作ページ</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/kokenLogo.ico" />
      </Head>

      <header>
        <h1 className={styles.title}>
          工研&times;鉄研プラレール展示 操作ページ
        </h1>
      </header>

      <main className={styles.main}>
        <section>
          <h2>映像部分</h2>
          <div
            style={{
              margin: 0,
              padding: 0,
              width: "100%",
              position: "relative",
            }}
          >
            {/* <VideoCast
              roomIds={roomIds}
              styles={[
                {
                  position: "relative",
                  zIndex: 1,
                  width: "100%",
                },
                {
                  position: "absolute",
                  zIndex: 2,
                  bottom: 4,
                  right: 0,
                  width: "25%",
                },
              ]}
            /> */}
          </div>
          <button
            onClick={() => {
              setRoomIds(["hachioji", "train1"])
            }}
          >
            京王八王子
          </button>
          <button
            onClick={() => {
              setRoomIds(["kitano", "train1"])
            }}
          >
            調布
          </button>
          <button
            onClick={() => {
              setRoomIds(["chofu", "train1"])
            }}
          >
            調布
          </button>
          <button
            onClick={() => {
              setRoomIds(["meidaimae", "train1"])
            }}
          >
            明大前
          </button>
          <button
            onClick={() => {
              setRoomIds(["sasazuka", "train1"])
            }}
          >
            笹塚
          </button>
          <button
            onClick={() => {
              setRoomIds(["iwamotocho", "train1"])
            }}
          >
            岩本町
          </button>
          <button
            onClick={() => {
              setRoomIds(["train1"])
            }}
          >
            車両前景
          </button>
        </section>

        <section>
          <h2>地図部分</h2>
          <RailroadMap
            datas={{
              stop: stopPointState,
              switchState: switchPointState,
              blockState: blockState,
              train1: {
                positionScale: trainPosition1,
                id: "koken",
              },
            }}
            onStopPointOrSwitchPointClick={(stationId) =>
              setSelectedStationId(stationId)
            }
          />
        </section>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://www.koken.club.uec.ac.jp/"
          target="_blank"
          rel="noopener noreferrer"
        >
          &copy;2022 電気通信大学工学研究部
        </a>
      </footer>
    </div>
  )
}

export default Home
