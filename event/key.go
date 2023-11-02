//go:generate go run github.com/dmarkham/enumer -type=Key

package event

// Key represents a Unity Bridge event key. Keys represent actions that
// and can be performed, and attributes that can be read or written.
// This is used to control the robot through the bridge.
type Key int32

const (
	KeyNone Key = iota
	KeyProductTest
	KeyProductType
	KeyCameraConnection
	KeyCameraFirmwareVersion
	KeyCameraStartShootPhoto
	KeyCameraIsShootingPhoto
	KeyCameraPhotoSize
	KeyCameraStartRecordVideo
	KeyCameraStopRecordVideo
	KeyCameraIsRecording
	KeyCameraCurrentRecordingTimeInSeconds
	KeyCameraVideoFormat
	KeyCameraMode
	KeyCameraDigitalZoomFactor
	KeyCameraAntiFlicker
	KeyCameraSwitch
	KeyCameraCurrentCameraIndex
	KeyCameraHasMainCamera
	KeyCameraHasSecondaryCamera
	KeyCameraFormatSDCard
	KeyCameraSDCardIsFormatting
	KeyCameraSDCardIsFull
	KeyCameraSDCardHasError
	KeyCameraSDCardIsInserted
	KeyCameraSDCardTotalSpaceInMB
	KeyCameraSDCardRemaingSpaceInMB
	KeyCameraSDCardAvailablePhotoCount
	KeyCameraSDCardAvailableRecordingTimeInSeconds
	KeyCameraIsTimeSynced
	KeyCameraDate
	KeyCameraVideoTransRate
	KeyCameraRequestIFrame
	KeyCameraAntiLarsenAlgorithmEnable
	KeyMainControllerConnection
	KeyMainControllerFirmwareVersion
	KeyMainControllerLoaderVersion
	KeyMainControllerVirtualStick
	KeyMainControllerVirtualStickEnabled
	KeyMainControllerChassisSpeedMode
	KeyMainControllerChassisFollowMode
	KeyMainControllerChassisCarControlMode
	KeyMainControllerRecordState
	KeyMainControllerGetRecordSetting
	KeyMainControllerSetRecordSetting
	KeyMainControllerPlayRecordAttr
	KeyMainControllerGetPlayRecordSetting
	KeyMainControllerSetPlayRecordSetting
	KeyMainControllerMaxSpeedForward
	KeyMainControllerMaxSpeedBackward
	KeyMainControllerMaxSpeedLateral
	KeyMainControllerSlopeY
	KeyMainControllerSlopeX
	KeyMainControllerSlopeBreakY
	KeyMainControllerSlopeBreakX
	KeyMainControllerMaxSpeedForwardConfig
	KeyMainControllerMaxSpeedBackwardConfig
	KeyMainControllerMaxSpeedLateralConfig
	KeyMainControllerSlopSpeedYConfig
	KeyMainControllerSlopSpeedXConfig
	KeyMainControllerSlopBreakYConfig
	KeyMainControllerSlopBreakXConfig
	KeyMainControllerChassisPosition
	KeyMainControllerWheelSpeed
	KeyRobomasterMainControllerEscEncodingStatus
	KeyRobomasterMainControllerEscEncodeFlag
	KeyRobomasterMainControllerStartIMUCalibration
	KeyRobomasterMainControllerIMUCalibrationState
	KeyRobomasterMainControllerIMUCalibrationCurrSide
	KeyRobomasterMainControllerIMUCalibrationProgress
	KeyRobomasterMainControllerIMUCalibrationFailCode
	KeyRobomasterMainControllerIMUCalibrationFinishFlag
	KeyRobomasterMainControllerStopIMUCalibration
	KeyRobomasterChassisMode
	KeyRobomasterChassisSpeed
	KeyRobomasterOpenChassisSpeedUpdates
	KeyRobomasterCloseChassisSpeedUpdates
	KeyRobomasterMainControllerRelativePosition
	KeyMainControllerArmServoID
	KeyMainControllerServoAddressing
	KeyRemoteControllerConnection
	KeyGimbalConnection
	KeyGimbalESCFirmwareVersion
	KeyGimbalFirmwareVersion
	KeyGimbalWorkMode
	KeyGimbalControlMode
	KeyGimbalResetPosition
	KeyGimbalResetPositionState
	KeyGimbalCalibration
	KeyGimbalSpeedRotation
	KeyGimbalSpeedRotationEnabled
	KeyGimbalAngleIncrementRotation
	KeyGimbalAngleFrontYawRotation
	KeyGimbalAngleFrontPitchRotation
	KeyGimbalAttitude
	KeyGimbalAutoCalibrate
	KeyGimbalCalibrationStatus
	KeyGimbalCalibrationProgress
	KeyGimbalOpenAttitudeUpdates
	KeyGimbalCloseAttitudeUpdates
	KeyRobomasterSystemConnection
	KeyRobomasterSystemFirmwareVersion
	KeyRobomasterSystemCANFirmwareVersion
	KeyRobomasterSystemScratchFirmwareVersion
	KeyRobomasterSystemSerialNumber
	KeyRobomasterSystemAbilitiesAttack
	KeyRobomasterSystemUnderAbilitiesAttack
	KeyRobomasterSystemKill
	KeyRobomasterSystemRevive
	KeyRobomasterSystemGet1860LinkAck
	KeyMainControllerGetLinkAck
	KeyGimbalGetLinkAck
	KeyRobomasterSystemGameRoleConfig
	KeyRobomasterSystemGameColorConfig
	KeyRobomasterSystemGameStart
	KeyRobomasterSystemGameEnd
	KeyRobomasterSystemDebugLog
	KeyRobomasterSystemSoundEnabled
	KeyRobomasterSystemLeftHeadlightBrightness
	KeyRobomasterSystemRightHeadlightBrightness
	KeyRobomasterSystemLEDColor
	KeyRobomasterSystemUploadScratch
	KeyRobomasterSystemUploadScratchByFTP
	KeyRobomasterSystemUninstallScratchSkill
	KeyRobomasterSystemInstallScratchSkill
	KeyRobomasterSystemInquiryDspMd5
	KeyRobomasterSystemInquiryDspMd5Ack
	KeyRobomasterSystemInquiryDspResourceMd5
	KeyRobomasterSystemInquiryDspResourceMd5Ack
	KeyRobomasterSystemLaunchSinglePlayerCustomSkill
	KeyRobomasterSystemStopSinglePlayerCustomSkill
	KeyRobomasterSystemControlScratch
	KeyRobomasterSystemScratchState
	KeyRobomasterSystemScratchCallback
	KeyRobomasterSystemForesightPosition
	KeyRobomasterSystemPullLogFiles
	KeyRobomasterSystemCurrentHP
	KeyRobomasterSystemTotalHP
	KeyRobomasterSystemCurrentBullets
	KeyRobomasterSystemTotalBullets
	KeyRobomasterSystemEquipments
	KeyRobomasterSystemBuffs
	KeyRobomasterSystemSkillStatus
	KeyRobomasterSystemGunCoolDown
	KeyRobomasterSystemGameConfigList
	KeyRobomasterSystemCarAndSkillID
	KeyRobomasterSystemAppStatus
	KeyRobomasterSystemLaunchMultiPlayerSkill
	KeyRobomasterSystemStopMultiPlayerSkill
	KeyRobomasterSystemConfigSkillTable
	KeyRobomasterSystemWorkingDevices
	KeyRobomasterSystemExceptions
	KeyRobomasterSystemTaskStatus
	KeyRobomasterSystemReturnEnabled
	KeyRobomasterSystemSafeMode
	KeyRobomasterSystemScratchExecuteState
	KeyRobomasterSystemAttitudeInfo
	KeyRobomasterSystemSightBeadPosition
	KeyRobomasterSystemSpeakerLanguage
	KeyRobomasterSystemSpeakerVolumn
	KeyRobomasterSystemChassisSpeedLevel
	KeyRobomasterSystemIsEncryptedFirmware
	KeyRobomasterSystemScratchErrorInfo
	KeyRobomasterSystemScratchOutputInfo
	KeyRobomasterSystemBarrelCoolDown
	KeyRobomasterSystemResetBarrelOverheat
	KeyRobomasterSystemMobileAccelerInfo
	KeyRobomasterSystemMobileGyroAttitudeAngleInfo
	KeyRobomasterSystemMobileGyroRotationRateInfo
	KeyRobomasterSystemEnableAcceleratorSubscribe
	KeyRobomasterSystemEnableGyroRotationRateSubscribe
	KeyRobomasterSystemEnableGyroAttitudeAngleSubscribe
	KeyRobomasterSystemDeactivate
	KeyRobomasterSystemFunctionEnable
	KeyRobomasterSystemIsGameRunning
	KeyRobomasterSystemIsActivated
	KeyRobomasterSystemLowPowerConsumption
	KeyRobomasterSystemEnterLowPowerConsumption
	KeyRobomasterSystemExitLowPowerConsumption
	KeyRobomasterSystemIsLowPowerConsumption
	KeyRobomasterSystemPushFile
	KeyRobomasterSystemPlaySound
	KeyRobomasterSystemPlaySoundStatus
	KeyRobomasterSystemCustomUIAttribute
	KeyRobomasterSystemCustomUIFunctionEvent
	KeyRobomasterSystemTotalMileage
	KeyRobomasterSystemTotalDrivingTime
	KeyRobomasterSystemSetPlayMode
	KeyRobomasterSystemCustomSkillInfo
	KeyRobomasterSystemAddressing
	KeyRobomasterSystemLEDLightEffect
	KeyRobomasterSystemOpenImageTransmission
	KeyRobomasterSystemCloseImageTransmission
	KeyVisionFirmwareVersion
	KeyVisionTrackingAutoLockTarget
	KeyVisionARParameters
	KeyVisionARTagEnabled
	KeyVisionDebugRect
	KeyVisionLaserPosition
	KeyVisionDetectionEnable
	KeyVisionMarkerRunningStatus
	KeyVisionTrackingRunningStatus
	KeyVisionAimbotRunningStatus
	KeyVisionHeadAndShoulderStatus
	KeyVisionHumanDetectionRunningStatus
	KeyVisionUserConfirm
	KeyVisionUserCancel
	KeyVisionUserTrackingRect
	KeyVisionTrackingDistance
	KeyVisionLineColor
	KeyVisionMarkerColor
	KeyVisionMarkerAdvanceStatus
	KeyPerceptionFirmwareVersion
	KeyPerceptionMarkerEnable
	KeyPerceptionMarkerResult
	KeyESCFirmwareVersion1
	KeyESCFirmwareVersion2
	KeyESCFirmwareVersion3
	KeyESCFirmwareVersion4
	KeyESCMotorInfomation1
	KeyESCMotorInfomation2
	KeyESCMotorInfomation3
	KeyESCMotorInfomation4
	KeyWiFiLinkFirmwareVersion
	KeyWiFiLinkDebugInfo
	KeyWiFiLinkMode
	KeyWiFiLinkSSID
	KeyWiFiLinkPassword
	KeyWiFiLinkAvailableChannelNumbers
	KeyWiFiLinkCurrentChannelNumber
	KeyWiFiLinkSNR
	KeyWiFiLinkSNRPushEnabled
	KeyWiFiLinkReboot
	KeyWiFiLinkChannelSelectionMode
	KeyWiFiLinkInterference
	KeyWiFiLinkDeleteNetworkConfig
	KeySDRLinkSNR
	KeySDRLinkBandwidth
	KeySDRLinkChannelSelectionMode
	KeySDRLinkCurrentFreqPoint
	KeySDRLinkCurrentFreqBand
	KeySDRLinkIsDualFreqSupported
	KeySDRLinkUpdateConfigs
	KeyAirLinkConnection
	KeyAirLinkSignalQuality
	KeyAirLinkCountryCode
	KeyAirLinkCountryCodeUpdated
	KeyArmorFirmwareVersion1
	KeyArmorFirmwareVersion2
	KeyArmorFirmwareVersion3
	KeyArmorFirmwareVersion4
	KeyArmorFirmwareVersion5
	KeyArmorFirmwareVersion6
	KeyArmorUnderAttack
	KeyArmorEnterResetID
	KeyArmorCancelResetID
	KeyArmorSkipCurrentID
	KeyArmorResetStatus
	KeyRobomasterWaterGunFirmwareVersion
	KeyRobomasterWaterGunWaterGunFire
	KeyRobomasterWaterGunWaterGunFireWithTimes
	KeyRobomasterWaterGunShootSpeed
	KeyRobomasterWaterGunShootFrequency
	KeyRobomasterInfraredGunConnection
	KeyRobomasterInfraredGunFirmwareVersion
	KeyRobomasterInfraredGunInfraredGunFire
	KeyRobomasterInfraredGunShootFrequency
	KeyRobomasterBatteryFirmwareVersion
	KeyRobomasterBatteryPowerPercent
	KeyRobomasterBatteryVoltage
	KeyRobomasterBatteryTemperature
	KeyRobomasterBatteryCurrent
	KeyRobomasterBatteryShutdown
	KeyRobomasterBatteryReboot
	KeyRobomasterGamePadConnection
	KeyRobomasterGamePadFirmwareVersion
	KeyRobomasterGamePadHasMouse
	KeyRobomasterGamePadHasKeyboard
	KeyRobomasterGamePadCtrlSensitivityX
	KeyRobomasterGamePadCtrlSensitivityY
	KeyRobomasterGamePadCtrlSensitivityYaw
	KeyRobomasterGamePadCtrlSensitivityYawSlop
	KeyRobomasterGamePadCtrlSensitivityYawDeadZone
	KeyRobomasterGamePadCtrlSensitivityPitch
	KeyRobomasterGamePadCtrlSensitivityPitchSlop
	KeyRobomasterGamePadCtrlSensitivityPitchDeadZone
	KeyRobomasterGamePadMouseLeftButton
	KeyRobomasterGamePadMouseRightButton
	KeyRobomasterGamePadC1
	KeyRobomasterGamePadC2
	KeyRobomasterGamePadFire
	KeyRobomasterGamePadFn
	KeyRobomasterGamePadNoCalibrate
	KeyRobomasterGamePadNotAtMiddle
	KeyRobomasterGamePadBatteryWarning
	KeyRobomasterGamePadBatteryPercent
	KeyRobomasterGamePadActivationSettings
	KeyRobomasterGamePadControlEnabled
	KeyRobomasterClawConnection
	KeyRobomasterClawFirmwareVersion
	KeyRobomasterClawCtrl
	KeyRobomasterClawStatus
	KeyRobomasterClawInfoSubscribe
	KeyRobomasterEnableClawInfoSubscribe
	KeyRobomasterArmConnection
	KeyRobomasterArmCtrl
	KeyRobomasterArmCtrlMode
	KeyRobomasterArmCalibration
	KeyRobomasterArmBlockedFlag
	KeyRobomasterArmPositionSubscribe
	KeyRobomasterArmReachLimitX
	KeyRobomasterArmReachLimitY
	KeyRobomasterEnableArmInfoSubscribe
	KeyRobomasterArmControlMode
	KeyRobomasterTOFConnection
	KeyRobomasterTOFLEDColor
	KeyRobomasterTOFOnlineModules
	KeyRobomasterTOFInfoSubscribe
	KeyRobomasterEnableTOFInfoSubscribe
	KeyRobomasterTOFFirmwareVersion1
	KeyRobomasterTOFFirmwareVersion2
	KeyRobomasterTOFFirmwareVersion3
	KeyRobomasterTOFFirmwareVersion4
	KeyRobomasterServoConnection
	KeyRobomasterServoLEDColor
	KeyRobomasterServoSpeed
	KeyRobomasterServoOnlineModules
	KeyRobomasterServoInfoSubscribe
	KeyRobomasterEnableServoInfoSubscribe
	KeyRobomasterServoFirmwareVersion1
	KeyRobomasterServoFirmwareVersion2
	KeyRobomasterServoFirmwareVersion3
	KeyRobomasterServoFirmwareVersion4
	KeyRobomasterSensorAdapterConnection
	KeyRobomasterSensorAdapterOnlineModules
	KeyRobomasterSensorAdapterInfoSubscribe
	KeyRobomasterEnableSensorAdapterInfoSubscribe
	KeyRobomasterSensorAdapterFirmwareVersion1
	KeyRobomasterSensorAdapterFirmwareVersion2
	KeyRobomasterSensorAdapterFirmwareVersion3
	KeyRobomasterSensorAdapterFirmwareVersion4
	KeyRobomasterSensorAdapterFirmwareVersion5
	KeyRobomasterSensorAdapterFirmwareVersion6
	KeyRobomasterSensorAdapterLEDColor
	KeyCount
)