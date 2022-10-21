import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import RailroadMap from "../components/RailRoadMap"
import VideoCast from "../components/VideoCast"
import { useEffect, useRef, useState } from "react"
import {
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

type SwitchPointState = Record<BunkiRailId, boolean>
const INITIAL_SWITCH_POINT_STATE: SwitchPointState = {
  iwamotocho_b1: false,
  iwamotocho_b2: false,
  iwamotocho_b3: false,
  iwamotocho_b4: false,
  sasazuka_b1: false,
  sasazuka_b2: false,
  chofu_b1: false,
  chofu_b2: false,
  chofu_b3: false,
  chofu_b4: false,
  chofu_b5: false,
  kitano_b1: false,
  kitano_b2: false,
  kitano_b3: false,
}

type SpeedState = Record<TrainId, number>
const INITIAL_SPEED_STATE: SpeedState = {
  TAKAO: 0,
  CHICHIBU: 0,
  HAKONE: 0,
  OKUTAMA: 0,
  NIKKO: 0,
  ENOSHIMA: 0,
  KAMAKURA: 0,
  YOKOSUKA: 0,
}
const INITIAL_SELECTED_TRAIN_ID: TrainId = "TAKAO"

const stationIdAndTextMap: Record<StationId, string> = {
  motoyawata_s1: "本八幡1番線",
  motoyawata_s2: "本八幡2番線",
  iwamotocho_s1: "岩本町1番線",
  iwamotocho_s2: "岩本町2, 3番線",
  iwamotocho_s4: "岩本町4番線",
  kudanshita_s5: "九段下1番線",
  kudanshita_s6: "九段下2番線",
  sasazuka_s1: "笹塚1番線",
  sasazuka_s2: "笹塚2番線",
  sasazuka_s3: "笹塚3番線",
  sasazuka_s4: "笹塚4番線",
  sasazuka_s5: "笹塚新線",
  meidaimae_s1: "明大前1番線",
  meidaimae_s2: "明大前2番線",
  chofu_s1: "調布1番線",
  chofu_s2: "調布2番線",
  chofu_s3: "調布3番線",
  chofu_s4: "調布4番線",
  chofu_s5: "調布-若葉台",
  chofu_s6: "unknown",
  kitano_s1: "北野1番線",
  kitano_s2: "北野2番線",
  kitano_s3: "北野3番線",
  kitano_s4: "北野4番線",
  kitano_s5: "北野-高尾",
  kitano_s6: "北野-京王八王子",
  kitano_s7: "unknown",
  takao_s1: "高尾",
  takao_s2: "高尾山口",
  iwamotocho_b1: "岩本町1-2,3番線分岐",
  iwamotocho_b2: "unknown",
  iwamotocho_b3: "unknown",
  iwamotocho_b4: "岩本町2,3-4番線分岐",
  sasazuka_b1: "笹塚新線-明大前分岐",
  sasazuka_b2: "笹塚3-4番線分岐",
  chofu_b1: "調布1-2番線分岐",
  chofu_b2: "調布2番線若葉台-北野分岐",
  chofu_b3: "若葉台-北野分岐",
  chofu_b4: "調布3-4番線分岐2",
  chofu_b5: "調布3-4番線分岐1",
  kitano_b1: "北野-高尾分岐",
  kitano_b2: "北野3-4番線分岐",
  kitano_b3: "unknown",
  unknown: "unknown",
}

const Home: NextPage = () => {
  const stationWs = useRef<WebSocket>()
  const [stopPointState, setStopPointState] = useState<StopPointState>(
    INITIAL_STOP_POINT_STATE
  )
  const [switchPointState, setSwitchPointState] = useState<SwitchPointState>(
    INITIAL_SWITCH_POINT_STATE
  )
  const [selectedStationId, setSelectedStationId] =
    useState<StationId>("unknown")
  const [trainPosition1, setTrainPosition1] = useState<number>(0.4)

  const speedWs = useRef<WebSocket>()
  const [speedState, setSpeedState] = useState<SpeedState>(INITIAL_SPEED_STATE)
  const [selectedTrainId, setSelectedTrainId] = useState<TrainId>(
    INITIAL_SELECTED_TRAIN_ID
  )
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
    const ws = new WebSocket("wss://speed.chofufes2022.gotti.dev/speed")
    speedWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
    })
    ws.addEventListener("message", (e) => {
      const message: SpeedMessage = JSON.parse(e.data)
      console.log(message)
      setSpeedState((previousState) => ({
        ...previousState,
        [message.train_name]: message.speed,
      }))
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
    const ws = new WebSocket("wss://control.chofufes2022.gotti.dev/ws")
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
            <VideoCast
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
            />
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
            北野
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

        <section>
          <h2>操作部分</h2>
          <p>
            選択中：
            {stationIdAndTextMap[selectedStationId]}
            <button
              type="button"
              onClick={() =>
                toggleStopPointOrSwitchPointState(selectedStationId)
              }
            >
              切り替え
            </button>
          </p>
          <svg width="100%" viewBox="0 0 200 100">
            <rect x={0} y={0} width={200} height={100} fill="dimgrey" />
            <SpeedMeter
              cx={80}
              cy={40}
              r={30}
              max={100}
              value={Math.abs(speedState[selectedTrainId])}
            />
            <ReverseHandle
              cx={150}
              cy={50}
              r={3}
              isBack={isBack}
              onChange={(nextIsBack) => {
                setIsBack(nextIsBack)
              }}
            />
          </svg>
          <input
            type="range"
            min={0}
            defaultValue={0}
            max={100}
            step={5}
            onChange={(e) => {
              const a = isBack ? -1 : 1 // 前進なら1、逆進なら-1になる係数
              const message: SpeedMessage = {
                train_name: selectedTrainId,
                speed: a * Number.parseInt(e.target.value),
              }
              speedWs.current?.send(JSON.stringify(message))
            }}
          />
        </section>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://www.koken.club.uec.ac.jp/"
          target="_blank"
          rel="noopener noreferrer"
        >
          &copy;2021 電気通信大学工学研究部
        </a>
      </footer>
    </div>
  )
}

export default Home
