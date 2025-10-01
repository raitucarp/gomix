package styles

func TransitionNormal() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionBehaviorProp): "normal",
		}
	}
}

func TransitionDiscrete() ApplyProp {
	return func(s *style) styleProp {
		return &properties{
			string(transitionBehaviorProp): "allow-discrete",
		}
	}
}
