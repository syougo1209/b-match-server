import { useState, useCallback } from 'react'
import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../../Screens'

type useRegistrationScreenReturn = {
  step: number
  onPressBackButton: ()=>void
  onPressSubmitButton: ()=>void
}

export const useRegistrationScreen = (): useRegistrationScreenReturn => {
  const [step, setStep] = useState<number>(0)
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Profile'>>()
  const onPressBackButton = step === 0 ? useCallback(() => navigation.navigate('Profile'), [step]) : useCallback(() => setStep((step)=> step - 1),[step])
  const onPressSubmitButton = useCallback(() => setStep((step)=> step + 1), [step])

  return {
    step,
    onPressBackButton,
    onPressSubmitButton,
  }
}
