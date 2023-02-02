import { useForm, UseFormReturn } from 'react-hook-form'
import { sexType } from '../../../../types/sexType'

type TUserRegistrationForm = {
  nickname: string
  mailAddress: string
  year: number
  month: number
  date: number
  password: string
  prefectureId: number
  sexType: sexType
}

type useRegisterFormReturn = {
  onSubmit: () => void
  methods: UseFormReturn<TUserRegistrationForm>
}
export const useRegisterForm =(): useRegisterFormReturn => {
  const methods = useForm<TUserRegistrationForm>()
  const postUserProfile = async (data: TUserRegistrationForm) => console.log(data)
  const onSubmit = methods.handleSubmit(postUserProfile)

  return {
    methods,
    onSubmit
  }
}
