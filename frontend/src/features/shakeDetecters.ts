interface ShakeDetectorInit {
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

export class ShakeDetector extends EventTarget {
	private _currentMoveAmount: number;
	private readonly _shakeDetectInterval: number;
	private readonly _accelerationThreshold: number;
	private _lastTime: number = new Date().getTime();

	/**
	 * 直近のシェイク検出時の移動量（メートル）
	 */
	get moveAmount(): number {
		return this._currentMoveAmount;
	}

	constructor(
		init: ShakeDetectorInit = {
			shakeDetectInterval: 100,
			accelerationThreshold: 2,
		}
	) {
		super();
		this._shakeDetectInterval = init.shakeDetectInterval;
		this._accelerationThreshold = init.accelerationThreshold;
		this._currentMoveAmount = 0;
	}

	public enableShakeDetector = async () => {
		const isDeviceMotionSupported = "DeviceMotionEvent" in window;
		if (!isDeviceMotionSupported) {
			alert("このデバイスはモーションセンサーをサポートしていません。");
			return;
		}

		const permissionGranted = await this._requestMotionPermission();
		if (permissionGranted) {
			window.addEventListener("devicemotion", this._onDeviceMotionHandler);
		} else {
			alert(
				"モーションセンサーのアクセス許可が得られなかったため、シェイク検出機能は利用できません。"
			);
		}

		this.dispatchEvent(new Event("load"));
		return;
	};

	private _requestMotionPermission = async (): Promise<boolean> => {
		// iOS13以降のデバイスでのみ、モーションセンサーのアクセス許可をリクエストする必要があるらしい。
		// @ts-ignore
		if (typeof DeviceMotionEvent.requestPermission === "function") {
			try {
				// @ts-ignore
				const permissionState = await DeviceMotionEvent.requestPermission();
				return permissionState === "granted";
			} catch (error) {
				console.error(error);
				alert(error);
				alert(
					"モーションセンサーのアクセス許可のリクエスト中にエラーが発生しました。"
				);
				return false;
			}
		} else {
			return true;
		}
	};

	private _onDeviceMotionHandler = (event: DeviceMotionEvent) => {
		const currentTime = new Date().getTime();
		const deltaTime = currentTime - this._lastTime;
		if (deltaTime < this._shakeDetectInterval) return;

		const moveAmount = this._calculateMoveAmount(event, deltaTime);
		this._currentMoveAmount = moveAmount;
		this._lastTime = currentTime;
		this.dispatchEvent(new Event("shake"));
	};

	private _calculateMoveAmount = (
		event: DeviceMotionEvent,
		timeDiffMS: number
	) => {
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
		if (maxAcceleration < this._accelerationThreshold) {
			return 0;
		}

		// x,y,z軸のうち最大の加速度を使って、等加速度運動だとみなし単位時間（timeThresholdMS）中にどれだけ移動したかの推測値を計算
		const deltaTimeSec = timeDiffMS / 1000;
		// 加速度から移動量を計算（速度atをtで積分）
		const moveAmountMeter = (maxAcceleration * deltaTimeSec ** 2) / 2;

		return moveAmountMeter;
	};
}
