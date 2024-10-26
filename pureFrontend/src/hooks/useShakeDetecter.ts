import { useState, useRef } from 'react';

interface ShakeDetectorProps {
  /**
   * シェイク検出の閾値となる時間（ミリ秒）
    */
  shakeDetectInterval: number;
  /**
   * シェイク検出の閾値となる加速度
   * 0~1程度の加速度は殆ど動かさなくても発生するため、2以上の加速度のみを対象とする
   */
  accelerationThreshold: number;
}

export function useShakeDetecter(
  { shakeDetectInterval,accelerationThreshold }: ShakeDetectorProps = {
    shakeDetectInterval: 100,
    accelerationThreshold: 2,
  }
) {
  const [totalMoveAmount, setTotalMoveAmount] = useState<number>(0);
  const [currentMoveAmount, setCurrentMoveAmount] = useState<number>(0);
  const lastTime = useRef(new Date().getTime());

  const enableShakeDetector = async () => {
    const isDeviceMotionSupported = 'DeviceMotionEvent' in window;
    if (!isDeviceMotionSupported) {
      alert('このデバイスはモーションセンサーをサポートしていません。');
      return;
    }

    const permissionGranted = await _requestMotionPermission();
    if (permissionGranted) {
      window.addEventListener('devicemotion', _onDeviceMotionHandler);
    } else {
      alert(
        'モーションセンサーのアクセス許可が得られなかったため、シェイク検出機能は利用できません。'
      );
    }
  };

  const _requestMotionPermission = async (): Promise<boolean> => {
    // iOS13以降のデバイスでのみ、モーションセンサーのアクセス許可をリクエストする必要があるらしい。
    // @ts-ignore
    if (typeof DeviceMotionEvent.requestPermission === 'function') {
      try {
        // @ts-ignore
        const permissionState = await DeviceMotionEvent.requestPermission();
        return permissionState === 'granted';
      } catch (error) {
        console.error(error);
        alert(error);
        alert('モーションセンサーのアクセス許可のリクエスト中にエラーが発生しました。');
        return false;
      }
    } else {
      return true;
    }
  };

  const _onDeviceMotionHandler = (event: DeviceMotionEvent) => {
    const currentTime = new Date().getTime();
    const deltaTime = currentTime - lastTime.current;
    if (deltaTime < shakeDetectInterval) return;

    const moveAmount = _calculateMoveAmount(event, deltaTime);
    setCurrentMoveAmount(moveAmount);
    setTotalMoveAmount((prev) => prev + moveAmount);
    lastTime.current = currentTime;
  };

  const _calculateMoveAmount = (event: DeviceMotionEvent, timeDiffMS: number) => {
    const acceleration = event.acceleration;
    if (!acceleration) {
      return 0;
    }

    // x,y,z軸の加速度のうち、最大のものを取得。かつ、閾値未満の場合は0を返す
    const maxAcceleration = Math.max(
      Math.abs(acceleration.x ?? 0),
      Math.abs(acceleration.y ?? 0),
      Math.abs(acceleration.z ?? 0)
    );
    if (maxAcceleration < accelerationThreshold) {
      return 0;
    }

    // x,y,z軸のうち最大の加速度を使って、等加速度運動だとみなし単位時間（timeThresholdMS）中にどれだけ移動したかの推測値を計算
    const deltaTimeSec = timeDiffMS / 1000;
    // 加速度から移動量を計算（速度atをtで積分）
    const moveAmountMeter = (maxAcceleration * deltaTimeSec ** 2) / 2;

    return moveAmountMeter;
  };

  return {
    /**
     * デバイスの移動量(メートル)の合計値。`shakeDetectInterval`毎に更新される。
     */
    totalMoveAmount,
    /**
     * 直近の移動量(メートル)。`shakeDetectInterval`毎に更新される。
     */
    currentMoveAmount,
    /**
     * デバイスのモーションセンサーに対するイベントリスナーを登録する。
     * iOS13以降のデバイスでは、モーションセンサーのアクセス許可をリクエストする。
     * リクエストが許可された場合、モーションセンサーのイベントリスナーを登録する。
     */
    enableShakeDetector,
  };
}
